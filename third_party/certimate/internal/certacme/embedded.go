package certacme

import (
	"context"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/certimate-go/certimate/internal/embeddedhooks"
	"github.com/go-acme/lego/v5/challenge"
	"github.com/go-acme/lego/v5/challenge/dns01"
	legolog "github.com/go-acme/lego/v5/log"
)

// Lego's resolver and CNAME switch are global. Original execution used separate
// processes; embedded execution serializes only ACME to preserve those options.
var embeddedACMEPermit = make(chan struct{}, 1)

func enterEmbeddedACME(ctx context.Context) (func(), error) {
	if !embeddedhooks.EnabledFor(ctx) {
		return func() {}, nil
	}
	select {
	case embeddedACMEPermit <- struct{}{}:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	previousLogger := legolog.Default()
	previousResolver := dns01.DefaultClient()
	previousCNAME, hadCNAME := os.LookupEnv("LEGO_DISABLE_CNAME_SUPPORT")
	// Provider diagnostics may contain HTTP bodies and credentials. Engine emits
	// safe node lifecycle events, so raw Lego logs must not reach the host logger.
	legolog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil)))
	return func() {
		legolog.SetDefault(previousLogger)
		dns01.SetDefaultClient(previousResolver)
		if hadCNAME {
			_ = os.Setenv("LEGO_DISABLE_CNAME_SUPPORT", previousCNAME)
		} else {
			_ = os.Unsetenv("LEGO_DISABLE_CNAME_SUPPORT")
		}
		<-embeddedACMEPermit
	}, nil
}

type authorizedChallenge struct {
	provider     challenge.Provider
	ctx          context.Context
	domains      []string
	credentialID string
	dns          bool
}

func (p *authorizedChallenge) check(domain, keyAuth string) error {
	if err := embeddedhooks.Check(p.ctx, embeddedhooks.Operation{Action: "certificate.issue", ResourceKind: "workflow", Domains: p.domains}); err != nil {
		return err
	}
	// HTTP-01/SSH also retains decrypted credentials between client creation
	// and challenge callbacks. Recheck that credential at every real operation.
	if p.credentialID != "" {
		if err := embeddedhooks.Check(p.ctx, embeddedhooks.Operation{Action: "credential.use", ResourceKind: "credential", ResourceID: p.credentialID}); err != nil {
			return err
		}
	}
	if p.dns {
		info := dns01.GetChallengeInfo(p.ctx, domain, keyAuth)
		return embeddedhooks.Check(p.ctx, embeddedhooks.Operation{Action: "dns.challenge", ResourceKind: "credential", ResourceID: p.credentialID, Domains: []string{strings.TrimSuffix(info.EffectiveFQDN, ".")}})
	}
	return nil
}
func (p *authorizedChallenge) Present(ctx context.Context, domain, token, keyAuth string) error {
	if err := p.check(domain, keyAuth); err != nil {
		return err
	}
	return p.provider.Present(ctx, domain, token, keyAuth)
}
func (p *authorizedChallenge) CleanUp(ctx context.Context, domain, token, keyAuth string) error {
	if err := p.check(domain, keyAuth); err != nil {
		return err
	}
	return p.provider.CleanUp(ctx, domain, token, keyAuth)
}
func (p *authorizedChallenge) Timeout() (time.Duration, time.Duration) {
	if timeout, ok := p.provider.(challenge.ProviderTimeout); ok {
		return timeout.Timeout()
	}
	return dns01.DefaultPropagationTimeout, dns01.DefaultPollingInterval
}
func guardChallenge(ctx context.Context, provider challenge.Provider, domains []string, credentialID string, dns bool) challenge.Provider {
	if !embeddedhooks.EnabledFor(ctx) {
		return provider
	}
	return &authorizedChallenge{provider: provider, ctx: ctx, domains: append([]string(nil), domains...), credentialID: credentialID, dns: dns}
}
