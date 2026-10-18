package certacme

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/certimate-go/certimate/internal/embeddedhooks"
)

type drainingAccountRepository struct{ entered, release chan struct{} }

func (r *drainingAccountRepository) GetByCAAndEmail(context.Context, string, string, string) (*ACMEAccount, error) {
	close(r.entered)
	<-r.release
	return &ACMEAccount{}, nil
}
func (*drainingAccountRepository) Save(context.Context, *ACMEAccount) (*ACMEAccount, error) {
	panic("unexpected account persistence")
}

func TestEmbeddedAccountCancellationDrainsRegistration(t *testing.T) {
	embeddedhooks.Configure(&embeddedhooks.Hooks{})
	defer embeddedhooks.Configure(nil)
	repo := &drainingAccountRepository{entered: make(chan struct{}), release: make(chan struct{})}
	ctx, cancel := context.WithCancel(embeddedhooks.Bind(context.Background()))
	defer cancel()
	done := make(chan error, 1)
	go func() {
		_, err := acquireACMEAccount(ctx, &ACMEConfig{CADirUrl: "https://acme.invalid/directory"}, "synthetic@example.invalid", &acmeRegistrationDeps{accountRepo: repo})
		done <- err
	}()
	<-repo.entered
	cancel()
	select {
	case err := <-done:
		t.Fatalf("registration worker escaped before draining: %v", err)
	case <-time.After(30 * time.Millisecond):
	}
	close(repo.release)
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("got %v, want context canceled", err)
		}
	case <-time.After(time.Second):
		t.Fatal("registration did not finish after its worker drained")
	}
}
