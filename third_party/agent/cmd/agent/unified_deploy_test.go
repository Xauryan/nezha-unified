package main

import (
	"context"
	"encoding/json"
	"github.com/nezhahq/agent/model"
	"github.com/nezhahq/agent/pkg/agentproto"
	pb "github.com/nezhahq/agent/proto"
	"testing"
)

func TestUnifiedCapabilityIndependentAndLocalDisableHonored(t *testing.T) {
	setTestRuntimeConfig(model.AgentConfig{DisableCommandExecute: true})
	var got *pb.TaskResult
	e := agentTaskExecution{parent: context.Background(), send: func(r *pb.TaskResult) error { got = r; return nil }}
	if e.dispatchUnified(&pb.Task{Id: 1, Type: model.TaskTypeKeepalive}) {
		t.Fatal("extension intercepted legacy task")
	}
	if !e.dispatchUnified(&pb.Task{Id: 2, Type: agentproto.TaskCapabilities}) || got == nil {
		t.Fatal("capability response missing")
	}
	var cap agentproto.Capabilities
	if json.Unmarshal([]byte(got.Data), &cap) != nil || cap.Protocol != 1 || cap.CertificateDeploy || !got.Successful {
		t.Fatal("disabled capability advertised")
	}
}
func TestUnifiedRejectsMalformedWithoutEchoingPayload(t *testing.T) {
	var got *pb.TaskResult
	e := agentTaskExecution{parent: context.Background(), send: func(r *pb.TaskResult) error { got = r; return nil }}
	e.runUnified(&pb.Task{Id: 3, Type: agentproto.TaskCall, Data: `{"synthetic_secret":"sensitive-example"}`})
	if got == nil || got.Successful {
		t.Fatal("malformed extension accepted")
	}
	var out agentproto.Response
	if json.Unmarshal([]byte(got.Data), &out) != nil || out.Code != "invalid" || !out.Final {
		t.Fatal("structured failure missing")
	}
}
