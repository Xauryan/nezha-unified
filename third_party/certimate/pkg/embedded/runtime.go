//go:build certimate_full

package embedded

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"strings"
	"time"

	"github.com/certimate-go/certimate/internal/app"
	"github.com/certimate-go/certimate/internal/certificate"
	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/internal/domain/dtos"
	"github.com/certimate-go/certimate/internal/embeddedhooks"
	"github.com/certimate-go/certimate/internal/repository"
	"github.com/certimate-go/certimate/internal/settings"
	"github.com/certimate-go/certimate/internal/workflow"
	legolog "github.com/go-acme/lego/v5/log"
	"github.com/pocketbase/pocketbase/core"
)

type runtimeState struct {
	service            *workflow.WorkflowService
	previousACMELogger *slog.Logger
}

func RuntimeAvailable() bool { return true }

func (e *Engine) StartRuntime(ctx context.Context, hooks RuntimeHooks) error {
	e.runtimeMu.Lock()
	defer e.runtimeMu.Unlock()
	if e.closing.Load() {
		return errors.New("Certimate storage is closing")
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if e.runtime != nil {
		return errors.New("Certimate runtime is already started")
	}
	if hooks.Authorize == nil || hooks.ResolveCredential == nil || hooks.OnRunCreated == nil {
		return errors.New("Certimate runtime requires host authorization, credential resolution and run binding")
	}
	if err := app.BindEmbedded(e.app); err != nil {
		return err
	}
	embeddedhooks.Configure(&hooks)
	previousACMELogger := legolog.Default()
	legolog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil)))
	cleanup := func(err error) error {
		e.app.Cron().Stop()
		embeddedhooks.Configure(nil)
		legolog.SetDefault(previousACMELogger)
		app.ReleaseEmbedded(e.app)
		return err
	}
	if err := settings.InitEmbedded(); err != nil {
		return cleanup(err)
	}
	svc := workflow.NewEmbeddedService()
	if err := svc.InitSchedule(ctx); err != nil {
		_ = svc.Shutdown(context.Background())
		return cleanup(err)
	}
	state := &runtimeState{service: svc, previousACMELogger: previousACMELogger}
	certificateSvc := certificate.NewCertificateService(repository.NewACMEAccountRepository(), repository.NewCertificateRepository())
	if err := e.app.Cron().Add("cleanupCertificateExpired", "0 0 * * *", func() {
		e.runtimeMu.Lock()
		defer e.runtimeMu.Unlock()
		if e.runtime == state && !e.closing.Load() {
			_ = certificateSvc.CleanupExpired(context.Background())
		}
	}); err != nil {
		_ = svc.Shutdown(context.Background())
		return cleanup(err)
	}
	e.runtime = state
	e.app.Store().Set("nezha_runtime_ready", true)
	e.app.Cron().SetTimezone(time.Local)
	e.app.Cron().Start()
	return nil
}

// StopRuntime stops new cron work, cancels all runs and waits for their final
// persistence. A timeout leaves the database and authorization hooks available so
// an in-flight operation cannot fall back to standalone execution or a closed DB.
func (e *Engine) StopRuntime(ctx context.Context) error {
	e.runtimeMu.Lock()
	defer e.runtimeMu.Unlock()
	if e.runtime == nil {
		return nil
	}
	e.app.Cron().Stop()
	for _, job := range e.app.Cron().Jobs() {
		if strings.HasPrefix(job.Id(), "workflow#") || job.Id() == "cleanupWorkflowHistoryRuns" || job.Id() == "cleanupCertificateExpired" {
			e.app.Cron().Remove(job.Id())
		}
	}
	e.app.Store().Remove("nezha_runtime_ready")
	if err := e.runtime.service.Shutdown(ctx); err != nil {
		return err
	}
	legolog.SetDefault(e.runtime.previousACMELogger)
	e.runtime = nil
	embeddedhooks.Configure(nil)
	app.ReleaseEmbedded(e.app)
	return nil
}

func (e *Engine) StartWorkflow(ctx context.Context, id string) (string, error) {
	e.runtimeMu.Lock()
	defer e.runtimeMu.Unlock()
	if e.closing.Load() || e.runtime == nil {
		return "", errors.New("Certimate runtime is not active")
	}
	result, err := e.runtime.service.StartRun(ctx, &dtos.WorkflowStartRunReq{WorkflowId: id, RunTrigger: domain.WorkflowTriggerTypeManual})
	if err != nil {
		return "", err
	}
	return result.RunId, nil
}
func (e *Engine) CancelWorkflow(ctx context.Context, workflowID, runID string) error {
	e.runtimeMu.Lock()
	defer e.runtimeMu.Unlock()
	if e.closing.Load() || e.runtime == nil {
		return errors.New("Certimate runtime is not active")
	}
	_, err := e.runtime.service.CancelRun(ctx, &dtos.WorkflowCancelRunReq{WorkflowId: workflowID, RunId: runID})
	return err
}

func (e *Engine) NativeWorkflowStats(ctx context.Context) (WorkflowStatistics, error) {
	e.runtimeMu.Lock()
	defer e.runtimeMu.Unlock()
	if e.closing.Load() || e.runtime == nil {
		return WorkflowStatistics{}, ErrRuntimeUnavailable
	}
	stats, err := e.runtime.service.GetStatistics(ctx)
	if err != nil {
		return WorkflowStatistics{}, err
	}
	return WorkflowStatistics{Concurrency: stats.Concurrency, PendingRunIDs: stats.PendingRunIds, ProcessingRunIDs: stats.ProcessingRunIds}, nil
}

// Called under the native storage read lock, after the transaction commits.
func (e *Engine) reconcileNativeWrite(ctx context.Context, collection string, record *core.Record, deleted bool) error {
	e.runtimeMu.Lock()
	defer e.runtimeMu.Unlock()
	if e.runtime == nil {
		return nil
	}
	switch collection {
	case "workflow":
		return e.runtime.service.ReconcileWorkflow(ctx, record.Id, deleted)
	case "settings":
		return settings.RefreshEmbedded(record, deleted)
	}
	return nil
}
