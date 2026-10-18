package workflow

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/pocketbase/dbx"

	"github.com/certimate-go/certimate/internal/app"
	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/internal/domain/dtos"
	"github.com/certimate-go/certimate/internal/embeddedhooks"
	"github.com/certimate-go/certimate/internal/repository"
	"github.com/certimate-go/certimate/internal/settings"
	"github.com/certimate-go/certimate/internal/workflow/dispatcher"
)

type WorkflowService struct {
	dispatcher dispatcher.WorkflowDispatcher
	mu         sync.Mutex
	active     bool

	workflowRepo    workflowRepository
	workflowRunRepo workflowRunRepository
}

func NewWorkflowService(workflowRepo workflowRepository, workflowRunRepo workflowRunRepository) *WorkflowService {
	srv := &WorkflowService{
		dispatcher: dispatcher.GetSingletonDispatcher(),

		workflowRepo:    workflowRepo,
		workflowRunRepo: workflowRunRepo,
	}
	return srv
}

// NewEmbeddedService deliberately avoids the upstream process singleton so a
// stopped embedded database can be reopened safely within the host process.
func NewEmbeddedService() *WorkflowService {
	return &WorkflowService{dispatcher: dispatcher.New(), workflowRepo: repository.NewWorkflowRepository(), workflowRunRepo: repository.NewWorkflowRunRepository()}
}

func (s *WorkflowService) InitSchedule(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.active {
		return fmt.Errorf("workflow runtime is already active")
	}
	if err := s.dispatcher.Bootup(ctx); err != nil {
		return err
	}
	s.active = true
	if err := app.GetScheduler().Add("cleanupWorkflowHistoryRuns", "0 0 * * *", func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		if s.active {
			_ = s.cleanupHistoryRuns(context.Background())
		}
	}); err != nil {
		return err
	}
	workflows, err := s.workflowRepo.ListEnabledScheduled(ctx)
	if err != nil {
		return err
	}
	for _, w := range workflows {
		if err := registerWorkflowJob(s, w.Id, w.TriggerCron); err != nil {
			return err
		}
	}
	return nil
}

// ReconcileWorkflow applies scheduler effects after the native service commits.
// Request-only PocketBase hooks are not involved in embedded mutations.
func (s *WorkflowService) ReconcileWorkflow(ctx context.Context, id string, deleted bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.active {
		return nil
	}
	scheduler := app.GetScheduler()
	if deleted {
		scheduler.Remove(buildPbJobKey(id))
		return nil
	}
	w, err := s.workflowRepo.GetById(ctx, id)
	if err != nil {
		return err
	}
	if !w.Enabled || w.Trigger != domain.WorkflowTriggerTypeScheduled {
		scheduler.Remove(buildPbJobKey(id))
		return nil
	}
	return registerWorkflowJob(s, id, w.TriggerCron)
}

func (s *WorkflowService) GetStatistics(ctx context.Context) (*dtos.WorkflowStatisticsResp, error) {
	stats := s.dispatcher.GetStatistics()
	return &dtos.WorkflowStatisticsResp{
		Concurrency:      stats.Concurrency,
		PendingRunIds:    stats.PendingRunIds,
		ProcessingRunIds: stats.ProcessingRunIds,
	}, nil
}

func (s *WorkflowService) StartRun(ctx context.Context, req *dtos.WorkflowStartRunReq) (*dtos.WorkflowStartRunResp, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.active {
		return nil, fmt.Errorf("workflow runtime is not active")
	}
	ctx = embeddedhooks.Bind(ctx)
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	workflow, err := s.workflowRepo.GetById(ctx, req.WorkflowId)
	if err != nil {
		return nil, err
	}

	if req.RunTrigger == domain.WorkflowTriggerTypeManual && (workflow.LastRunStatus == domain.WorkflowRunStatusTypePending || workflow.LastRunStatus == domain.WorkflowRunStatusTypeProcessing) {
		return nil, fmt.Errorf("workflow is already pending or processing")
	} else if workflow.GraphContent == nil {
		return nil, fmt.Errorf("workflow graph content is empty")
	} else if err := workflow.GraphContent.Verify(); err != nil {
		return nil, fmt.Errorf("workflow graph content is invalid: %w", err)
	}

	if embeddedhooks.EnabledFor(ctx) {
		record, err := app.GetApp().FindRecordById(domain.CollectionNameWorkflow, workflow.Id)
		if err != nil || !workflow.Enabled || (record.GetString("unified_operation_id") != "" && !record.GetBool("unified_active")) {
			return nil, embeddedhooks.ErrDenied
		}
	}
	digest, err := embeddedhooks.GraphDigest(workflow.GraphContent)
	if err != nil {
		return nil, err
	}
	if err := embeddedhooks.Check(ctx, embeddedhooks.Operation{WorkflowID: workflow.Id, GraphDigest: digest, Action: "workflow.run", ResourceKind: "workflow", ResourceID: workflow.Id}); err != nil {
		return nil, err
	}

	workflowRun := &domain.WorkflowRun{
		WorkflowId: workflow.Id,
		Status:     domain.WorkflowRunStatusTypePending,
		Trigger:    req.RunTrigger,
		StartedAt:  time.Now(),
		Graph:      workflow.GraphContent.Clone(),
	}
	if resp, err := s.workflowRunRepo.SaveWithCascading(ctx, workflowRun); err != nil {
		return nil, err
	} else {
		workflowRun = resp
	}

	if err := embeddedhooks.RunCreated(ctx, workflow.Id, workflowRun.Id); err != nil {
		workflowRun.Status = domain.WorkflowRunStatusTypeFailed
		workflowRun.EndedAt = time.Now()
		workflowRun.Error = "embedded workflow authorization denied"
		_, saveErr := s.workflowRunRepo.SaveWithCascading(context.Background(), workflowRun)
		return nil, errors.Join(err, saveErr)
	}
	if err := s.dispatcher.Start(ctx, workflowRun.Id); err != nil {
		workflowRun.Status = domain.WorkflowRunStatusTypeFailed
		workflowRun.EndedAt = time.Now()
		workflowRun.Error = "workflow could not be queued"
		_, saveErr := s.workflowRunRepo.SaveWithCascading(context.Background(), workflowRun)
		return nil, errors.Join(err, saveErr)
	}

	return &dtos.WorkflowStartRunResp{RunId: workflowRun.Id}, nil
}

func (s *WorkflowService) CancelRun(ctx context.Context, req *dtos.WorkflowCancelRunReq) (*dtos.WorkflowCancelRunResp, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.active {
		return nil, fmt.Errorf("workflow runtime is not active")
	}
	workflow, err := s.workflowRepo.GetById(ctx, req.WorkflowId)
	if err != nil {
		return nil, err
	}

	workflowRun, err := s.workflowRunRepo.GetById(ctx, req.RunId)
	if err != nil {
		return nil, err
	} else if workflowRun.WorkflowId != workflow.Id {
		return nil, fmt.Errorf("workflow run not found")
	} else if workflowRun.Status != domain.WorkflowRunStatusTypePending && workflowRun.Status != domain.WorkflowRunStatusTypeProcessing {
		return nil, fmt.Errorf("workflow run is not pending or processing")
	}

	if err := s.dispatcher.Cancel(ctx, workflowRun.Id); err != nil {
		return nil, err
	}

	return &dtos.WorkflowCancelRunResp{}, nil
}

func (s *WorkflowService) Shutdown(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.active = false
	return s.dispatcher.Shutdown(ctx)
}

func (s *WorkflowService) cleanupHistoryRuns(ctx context.Context) error {
	globalSettingsForPersistence := settings.GetGlobalSettingsForPersistence()
	if globalSettingsForPersistence.WorkflowRunsRetentionMaxDays != 0 {
		ret, err := s.workflowRunRepo.DeleteWithExprs(ctx,
			dbx.NewExp(fmt.Sprintf("status!='%s'", domain.WorkflowRunStatusTypePending)),
			dbx.NewExp(fmt.Sprintf("status!='%s'", domain.WorkflowRunStatusTypeProcessing)),
			dbx.NewExp(fmt.Sprintf("endedAt<DATETIME('now', '-%d days')", globalSettingsForPersistence.WorkflowRunsRetentionMaxDays)),
		)
		if err != nil {
			app.GetLogger().Error("failed to delete workflow history runs", slog.Any("error", err))
			return err
		}

		if ret > 0 {
			app.GetLogger().Info(fmt.Sprintf("cleanup %d workflow history runs", ret))
		}
	}

	return nil
}
