package certacme

import (
	"context"
	"errors"
	"github.com/certimate-go/certimate/internal/embeddedhooks"
	"github.com/go-acme/lego/v5/challenge"
	"testing"
	"time"
)

type fakeChallenge struct{ present, cleanup int }

func (f *fakeChallenge) Present(context.Context, string, string, string) error {
	f.present++
	return nil
}
func (f *fakeChallenge) CleanUp(context.Context, string, string, string) error {
	f.cleanup++
	return nil
}
func (f *fakeChallenge) Timeout() (time.Duration, time.Duration) {
	return 7 * time.Minute, 3 * time.Second
}

func TestChallengeChecksRevocationAtProviderBoundary(t *testing.T) {
	revoked := false
	embeddedhooks.Configure(&embeddedhooks.Hooks{Authorize: func(_ context.Context, op embeddedhooks.Operation) error {
		if revoked {
			return errors.New("revoked")
		}
		if op.Action == "certificate.issue" && (len(op.Domains) != 2 || op.Domains[1] != "*.example.test") {
			t.Fatal("all requested SANs required")
		}
		return nil
	}})
	t.Cleanup(func() { embeddedhooks.Configure(nil) })
	upstream := &fakeChallenge{}
	p := guardChallenge(embeddedhooks.Bind(context.Background()), upstream, []string{"example.test", "*.example.test"}, "credential", false)
	if err := p.Present(context.Background(), "example.test", "synthetic", "synthetic"); err != nil {
		t.Fatal(err)
	}
	revoked = true
	if err := p.Present(context.Background(), "example.test", "synthetic", "synthetic"); !errors.Is(err, embeddedhooks.ErrDenied) {
		t.Fatal("revoked provider present was permitted")
	}
	if err := p.CleanUp(context.Background(), "example.test", "synthetic", "synthetic"); !errors.Is(err, embeddedhooks.ErrDenied) {
		t.Fatal("revoked provider cleanup was permitted")
	}
	if upstream.present != 1 || upstream.cleanup != 0 {
		t.Fatal("side effects happened after revocation")
	}
	timeout, interval := p.(challenge.ProviderTimeout).Timeout()
	if timeout != 7*time.Minute || interval != 3*time.Second {
		t.Fatal("upstream provider timeout lost")
	}
}

func TestHTTPChallengeRechecksCredentialBeforeEveryEffect(t *testing.T) {
	for _, reason := range []string{"revoked", "disabled", "rotated"} {
		t.Run(reason, func(t *testing.T) {
			blocked := false
			var credentialChecks int
			embeddedhooks.Configure(&embeddedhooks.Hooks{Authorize: func(_ context.Context, op embeddedhooks.Operation) error {
				if op.Action == "credential.use" {
					credentialChecks++
					if op.ResourceID != "synthetic-ssh-credential" || blocked {
						return errors.New(reason)
					}
				}
				return nil
			}})
			defer embeddedhooks.Configure(nil)
			upstream := &fakeChallenge{}
			p := guardChallenge(embeddedhooks.Bind(context.Background()), upstream, []string{"example.test"}, "synthetic-ssh-credential", false)
			if err := p.Present(context.Background(), "example.test", "synthetic", "synthetic"); err != nil {
				t.Fatal(err)
			}
			blocked = true
			if err := p.Present(context.Background(), "example.test", "synthetic", "synthetic"); !errors.Is(err, embeddedhooks.ErrDenied) {
				t.Fatal("HTTP challenge used a blocked credential")
			}
			if err := p.CleanUp(context.Background(), "example.test", "synthetic", "synthetic"); !errors.Is(err, embeddedhooks.ErrDenied) {
				t.Fatal("HTTP cleanup used a blocked credential")
			}
			if credentialChecks != 3 || upstream.present != 1 || upstream.cleanup != 0 {
				t.Fatalf("checks=%d, present=%d, cleanup=%d", credentialChecks, upstream.present, upstream.cleanup)
			}
		})
	}
}
