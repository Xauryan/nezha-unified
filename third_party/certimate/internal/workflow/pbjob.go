package workflow

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/certimate-go/certimate/internal/app"
	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/internal/domain/dtos"
)

func registerWorkflowJob(workflowSrv *WorkflowService, workflowId string, triggerCron string) error {
	scheduler := app.GetScheduler()

	jobKey := buildPbJobKey(workflowId)
	// Always replace the callback: the same embedded database can restart with a
	// fresh service while preserving the job expression.
	err := scheduler.Add(jobKey, triggerCron, func() {
		// StartRun checks lifecycle state under its mutex before accessing the app.
		// A cron callback already queued during shutdown must not touch a released
		// global application or log through its closed database.
		_, _ = workflowSrv.StartRun(context.Background(), &dtos.WorkflowStartRunReq{
			WorkflowId: workflowId,
			RunTrigger: domain.WorkflowTriggerTypeScheduled,
		})
	})
	if err != nil {
		app.GetLogger().Error(fmt.Sprintf("failed to register cron job for workflow #%s", workflowId), slog.Any("error", err))
		return fmt.Errorf("failed to add cron job: %w", err)
	}

	app.GetLogger().Info(fmt.Sprintf("registered cron job for workflow #%s", workflowId), slog.String("cron", triggerCron))
	return nil
}

func buildPbJobKey(workflowId string) string {
	return fmt.Sprintf("workflow#%s", workflowId)
}
