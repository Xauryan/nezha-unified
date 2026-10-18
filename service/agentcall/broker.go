// Package agentcall separates authenticated Agent transport from each caller's
// authorization, capability and cancellation policy. It does not depend on MCP.
package agentcall

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/nezhahq/nezha/pkg/agentproto"
	pb "github.com/nezhahq/nezha/proto"
)

var (
	ErrOffline  = errors.New("agent offline")
	ErrTimeout  = errors.New("agent result unknown after timeout")
	ErrDenied   = errors.New("agent call no longer authorized")
	ErrCapacity = errors.New("agent call capacity reached")
	ErrProtocol = errors.New("invalid agent response")
)

type Transport func(context.Context, uint64, *pb.Task) error
type Authorize func(context.Context) error
type Broker struct {
	mu      sync.Mutex
	pending map[uint64]*pending
	send    Transport
	limit   int
}
type pending struct {
	server, kind       uint64
	operation, purpose string
	results            chan *pb.TaskResult
	cancel             context.CancelFunc
}

func New(send Transport, limit int) *Broker {
	if limit < 1 {
		limit = 64
	}
	return &Broker{pending: map[uint64]*pending{}, send: send, limit: limit}
}

// Deliver must be called only after authentication AND current stream ownership
// checks. Both reporter and task type are bound to the registered request.
func (b *Broker) Deliver(server uint64, r *pb.TaskResult) bool {
	if r == nil || len(r.Data) > agentproto.MaxPayload {
		return false
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	p := b.pending[r.Id]
	if p == nil || server == 0 || p.server != server || p.kind != r.Type {
		return false
	}
	select {
	case p.results <- r:
		return true
	default:
		return false
	}
}
func (b *Broker) CancelPurpose(purpose string) int {
	b.mu.Lock()
	defer b.mu.Unlock()
	n := 0
	for _, p := range b.pending {
		if p.purpose == purpose {
			p.cancel()
			n++
		}
	}
	return n
}
func (b *Broker) Call(ctx context.Context, server, kind uint64, purpose string, in any, auth Authorize, progress func(agentproto.Response)) (json.RawMessage, error) {
	if auth == nil || auth(ctx) != nil {
		return nil, ErrDenied
	}
	if server == 0 || b.send == nil {
		return nil, ErrOffline
	}
	body, err := json.Marshal(in)
	if err != nil || len(body) > agentproto.MaxPayload {
		return nil, ErrProtocol
	}
	var operation string
	if req, ok := in.(agentproto.Request); ok {
		operation = req.OperationID
	}
	wait, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()
	p := &pending{server: server, kind: kind, purpose: purpose, operation: operation, results: make(chan *pb.TaskResult, 32), cancel: cancel}
	b.mu.Lock()
	if len(b.pending) >= b.limit {
		b.mu.Unlock()
		return nil, ErrCapacity
	}
	var id uint64
	for id == 0 || b.pending[id] != nil {
		var buf [8]byte
		if _, err = rand.Read(buf[:]); err != nil {
			b.mu.Unlock()
			return nil, err
		}
		id = (binary.BigEndian.Uint64(buf[:]) & ((1 << 52) - 1)) | (1 << 52)
	}
	b.pending[id] = p
	b.mu.Unlock()
	defer func() { b.mu.Lock(); delete(b.pending, id); b.mu.Unlock() }()
	if auth(wait) != nil {
		return nil, ErrDenied
	}
	if err = b.send(wait, server, &pb.Task{Id: id, Type: kind, Data: string(body)}); err != nil {
		return nil, err
	}
	for {
		select {
		case r := <-p.results:
			if wait.Err() != nil {
				return nil, wait.Err()
			}
			if auth(wait) != nil {
				return nil, ErrDenied
			}
			if kind == agentproto.TaskCapabilities {
				if !r.Successful {
					return nil, ErrProtocol
				}
				return json.RawMessage(r.Data), nil
			}
			var out agentproto.Response
			if json.Unmarshal([]byte(r.Data), &out) != nil || out.Protocol != agentproto.Protocol || out.OperationID != p.operation {
				return nil, ErrProtocol
			}
			if progress != nil {
				progress(out)
			}
			if out.Final {
				if !r.Successful {
					return out.Payload, errors.New("agent deployment failed: " + safeCode(out.Code))
				}
				return out.Payload, nil
			}
		case <-wait.Done():
			if errors.Is(wait.Err(), context.DeadlineExceeded) {
				return nil, ErrTimeout
			}
			return nil, wait.Err()
		}
	}
}
func safeCode(s string) string {
	switch s {
	case "invalid", "denied", "failed", "unknown", "stale", "conflict", "not_found":
		return s
	default:
		return "failed"
	}
}
func (b *Broker) Capabilities(ctx context.Context, server uint64, auth Authorize) (agentproto.Capabilities, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	var cap agentproto.Capabilities
	raw, err := b.Call(ctx, server, agentproto.TaskCapabilities, "capabilities", agentproto.Capabilities{Protocol: agentproto.Protocol}, auth, nil)
	if err != nil {
		return cap, err
	}
	if json.Unmarshal(raw, &cap) != nil || cap.Protocol != agentproto.Protocol {
		return cap, ErrProtocol
	}
	cap.Online = true
	return cap, nil
}
