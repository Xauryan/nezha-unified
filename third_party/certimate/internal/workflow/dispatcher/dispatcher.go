package dispatcher

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"runtime"
	"runtime/debug"
	"sync"
	"time"

	"github.com/certimate-go/certimate/internal/app"
	"github.com/certimate-go/certimate/internal/domain"
	"github.com/certimate-go/certimate/internal/embeddedhooks"
	"github.com/certimate-go/certimate/internal/repository"
	"github.com/certimate-go/certimate/internal/workflow/engine"
	"github.com/certimate-go/certimate/pkg/logging"
	xenv "github.com/certimate-go/certimate/pkg/utils/env"
)

var envMaxWorkers = 1

func init() {
	envMaxWorkers = xenv.GetOrDefaultInt("CERTIMATE_WORKFLOW_MAX_WORKERS", runtime.GOMAXPROCS(0))
	if envMaxWorkers <= 0 {
		envMaxWorkers = max(1, runtime.NumCPU())
	}
}

type WorkflowDispatcher interface {
	GetStatistics() Statistics

	Bootup(ctx context.Context) error
	Shutdown(ctx context.Context) error
	Start(ctx context.Context, runId string) error
	Cancel(ctx context.Context, runId string) error
}

type Statistics struct {
	Concurrency      int
	PendingRunIds    []string
	ProcessingRunIds []string
}

type workflowDispatcher struct {
	booted       bool
	concurrency  int
	taskWG       sync.WaitGroup
	shutdownDone chan struct{}

	taskMtx         sync.RWMutex
	pendingRunQueue []string
	processingTasks map[string]*taskInfo // Key: RunId

	workflowRepo    workflowRepository
	workflowRunRepo workflowRunRepository
	workflowLogRepo workflowLogRepository

	syslog *slog.Logger
}

var _ WorkflowDispatcher = (*workflowDispatcher)(nil)

func (wd *workflowDispatcher) GetStatistics() Statistics {
	wd.taskMtx.RLock()
	defer wd.taskMtx.RUnlock()

	stats := Statistics{
		Concurrency:      wd.concurrency,
		PendingRunIds:    make([]string, 0),
		ProcessingRunIds: make([]string, 0),
	}
	for _, pendingRunId := range wd.pendingRunQueue {
		stats.PendingRunIds = append(stats.PendingRunIds, pendingRunId)
	}
	for _, processingRunId := range wd.processingTasks {
		stats.ProcessingRunIds = append(stats.ProcessingRunIds, processingRunId.RunId)
	}

	return stats
}

func (wd *workflowDispatcher) Bootup(ctx context.Context) error {
	wd.taskMtx.Lock()
	defer wd.taskMtx.Unlock()
	if wd.booted {
		return fmt.Errorf("dispatcher is already running")
	}
	if wd.shutdownDone != nil {
		select {
		case <-wd.shutdownDone:
		default:
			return fmt.Errorf("dispatcher is still stopping")
		}
	}
	if err := wd.workflowRunRepo.ResetStatusIfHanging(ctx); err != nil {
		return err
	}
	wd.shutdownDone = make(chan struct{})
	wd.booted = true
	return nil
}

func (wd *workflowDispatcher) Shutdown(ctx context.Context) error {
	wd.taskMtx.Lock()
	done := wd.shutdownDone
	if wd.booted {
		wd.booted = false
		for _, task := range wd.processingTasks {
			task.cancel()
		}
		for _, id := range wd.pendingRunQueue {
			if run, err := wd.workflowRunRepo.GetById(context.Background(), id); err == nil {
				run.Status = domain.WorkflowRunStatusTypeCanceled
				run.EndedAt = time.Now()
				_, _ = wd.workflowRunRepo.SaveWithCascading(context.Background(), run)
			}
		}
		wd.pendingRunQueue = nil
		// No Add can occur once booted is false under taskMtx.
		go func() { wd.taskWG.Wait(); close(done) }()
	}
	wd.taskMtx.Unlock()
	if done == nil {
		return nil
	}
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (wd *workflowDispatcher) Start(ctx context.Context, runId string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	wd.taskMtx.Lock()
	defer wd.taskMtx.Unlock()
	if !wd.booted {
		return fmt.Errorf("dispatcher is not running")
	}
	if _, exists := wd.processingTasks[runId]; exists {
		return fmt.Errorf("workflow run is already processing")
	}
	for _, id := range wd.pendingRunQueue {
		if id == runId {
			return fmt.Errorf("workflow run is already queued")
		}
	}
	wd.pendingRunQueue = append(wd.pendingRunQueue, runId)
	wd.tryNextLocked()
	return nil
}

func (wd *workflowDispatcher) Cancel(ctx context.Context, runId string) error {
	wd.taskMtx.Lock()
	defer wd.taskMtx.Unlock()

	workflowRun, err := wd.workflowRunRepo.GetById(ctx, runId)
	if err != nil {
		return err
	} else if workflowRun.Status != domain.WorkflowRunStatusTypePending && workflowRun.Status != domain.WorkflowRunStatusTypeProcessing {
		return fmt.Errorf("workrun #%s is already completed", workflowRun.Id)
	}

	workflow, err := wd.workflowRepo.GetById(ctx, workflowRun.WorkflowId)
	if err != nil {
		return err
	}

	workflowRun.Status = domain.WorkflowRunStatusTypeCanceled
	workflowRun.EndedAt = time.Now()
	if workflow.LastRunId == workflowRun.Id {
		_, err := wd.workflowRunRepo.SaveWithCascading(ctx, workflowRun)
		if err != nil {
			return err
		}
	} else {
		_, err := wd.workflowRunRepo.Save(ctx, workflowRun)
		if err != nil {
			return err
		}
	}

	if task, exists := wd.processingTasks[runId]; exists {
		task.cancel()

		wd.syslog.Info(fmt.Sprintf("workrun #%s was canceled", task.RunId))
	}

	for i, pendingRunId := range wd.pendingRunQueue {
		if pendingRunId == runId {
			wd.pendingRunQueue = append(wd.pendingRunQueue[:i], wd.pendingRunQueue[i+1:]...)
			break
		}
	}

	wd.tryNextLocked()

	return nil
}

func (wd *workflowDispatcher) tryExecuteAsync(task *taskInfo) {
	defer wd.taskWG.Done()
	var workflow *domain.Workflow
	var workflowRun *domain.WorkflowRun
	var err error

	// 尝试继续执行等待队列中的任务
	defer func() {
		wd.taskMtx.Lock()
		delete(wd.processingTasks, task.RunId)
		wd.tryNextLocked()
		wd.taskMtx.Unlock()
	}()

	// 捕获 panic
	defer func() {
		if r := recover(); r != nil {
			if embeddedhooks.EnabledFor(task.ctx) {
				wd.syslog.Error("workflow dispatcher panic", slog.String("workflowId", task.WorkflowId), slog.String("runId", task.RunId))
			} else {
				wd.syslog.Error(fmt.Sprintf("workflow dispatcher panic: %v", r), slog.String("workflowId", task.WorkflowId), slog.String("runId", task.RunId))
				slog.Error(fmt.Sprintf("workflow dispatcher panic: %v, stack trace: %s", r, string(debug.Stack())), slog.String("workflowId", task.WorkflowId), slog.String("runId", task.RunId))
			}

			if workflowRun != nil {
				workflowRun.Status = domain.WorkflowRunStatusTypeFailed
				workflowRun.EndedAt = time.Now()
				workflowRun.Error = embeddedhooks.SafeError(task.ctx, fmt.Errorf("workflow dispatcher panic: %v", r)).Error()
				if _, err := wd.workflowRunRepo.SaveWithCascading(context.Background(), workflowRun); err != nil {
					log.Default().Println("failed to save workflow run after panic", slog.Any("error", err))
				}
			}
		}
	}()

	// 查询运行实体，并级联更新状态
	if workflowRun, err = wd.workflowRunRepo.GetById(task.ctx, task.RunId); err != nil {
		if !(errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded)) {
			wd.syslog.Error(fmt.Sprintf("failed to get workrun #%s record", task.RunId), slog.Any("error", err))
		}
		return
	} else {
		if workflowRun.Status == domain.WorkflowRunStatusTypePending {
			workflowRun.Status = domain.WorkflowRunStatusTypeProcessing
			wd.workflowRunRepo.SaveWithCascading(context.Background(), workflowRun)
		} else {
			// WTF? That should be impossible!
			return
		}
	}

	// 查询工作流实体
	workflow, err = wd.workflowRepo.GetById(task.ctx, workflowRun.WorkflowId)
	if err != nil {
		wd.syslog.Error(fmt.Sprintf("failed to get workflow #%s record", workflowRun.WorkflowId), slog.Any("error", err))
		return
	}

	// 初始化工作流引擎
	logsBuf := make(domain.WorkflowLogs, 0)
	var logsMu sync.Mutex
	we := engine.NewWorkflowEngine()
	terminal := false
	we.OnEnd(func(ctx context.Context) error {
		terminal = true
		logsMu.Lock()
		errmsg := logsBuf.ErrorString()
		logsMu.Unlock()
		if task.ctx.Err() != nil {
			workflowRun.Status = domain.WorkflowRunStatusTypeCanceled
			workflowRun.EndedAt = time.Now()
		} else if errmsg == "" {
			workflowRun.Status = domain.WorkflowRunStatusTypeSucceeded
			workflowRun.EndedAt = time.Now()
		} else {
			workflowRun.Status = domain.WorkflowRunStatusTypeFailed
			workflowRun.EndedAt = time.Now()
			workflowRun.Error = errmsg
		}
		wd.workflowRunRepo.SaveWithCascading(context.Background(), workflowRun)

		return nil
	})
	we.OnError(func(ctx context.Context, err error) error {
		terminal = true
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			workflowRun.Status = domain.WorkflowRunStatusTypeCanceled
			workflowRun.EndedAt = time.Now()
			wd.workflowRunRepo.SaveWithCascading(context.Background(), workflowRun)
		} else {
			workflowRun.Status = domain.WorkflowRunStatusTypeFailed
			workflowRun.EndedAt = time.Now()
			workflowRun.Error = err.Error()
			wd.workflowRunRepo.SaveWithCascading(context.Background(), workflowRun)
		}

		return nil
	})
	we.OnNodeError(func(ctx context.Context, node *engine.Node, err error) error {
		if errors.Is(err, engine.ErrTerminated) || errors.Is(err, engine.ErrBlocksException) {
			return nil
		}

		log := domain.WorkflowLog{}
		log.WorkflowId = task.WorkflowId
		log.RunId = task.RunId
		log.NodeId = node.Id
		log.NodeName = node.Data.Name
		log.TimestampMilli = time.Now().UnixMilli()
		log.Level = int32(slog.LevelError)
		log.Message = err.Error()
		log.CreatedAt = time.Now()
		logsMu.Lock()
		logsBuf = append(logsBuf, log)
		logsMu.Unlock()

		if _, err := wd.workflowLogRepo.Save(ctx, &log); err != nil {
			wd.syslog.Error(err.Error())
		}

		return nil
	})
	we.OnNodeLogging(func(ctx context.Context, node *engine.Node, record logging.Record) error {
		log := domain.WorkflowLog{}
		log.WorkflowId = task.WorkflowId
		log.RunId = task.RunId
		log.NodeId = node.Id
		log.NodeName = node.Data.Name
		log.TimestampMilli = record.Time.UnixMilli()
		log.Level = int32(record.Level)
		log.Message = record.Message
		log.Data = record.Data()
		log.CreatedAt = time.Now()
		logsMu.Lock()
		logsBuf = append(logsBuf, log)
		logsMu.Unlock()

		if _, err := wd.workflowLogRepo.Save(ctx, &log); err != nil {
			wd.syslog.Error(err.Error())
		}

		return nil
	})

	// 执行工作流
	wd.syslog.Info(fmt.Sprintf("workflow #%s's run #%s started", task.WorkflowId, task.RunId))
	invokeErr := we.Invoke(task.ctx, engine.WorkflowExecution{
		WorkflowId:          workflow.Id,
		WorkflowName:        workflow.Name,
		WorkflowDescription: workflow.Description,
		RunId:               workflowRun.Id,
		RunTrigger:          workflowRun.Trigger,
		RunAt:               workflowRun.StartedAt,
		Graph:               workflowRun.Graph,
	})
	if !terminal {
		workflowRun.Status = domain.WorkflowRunStatusTypeFailed
		if task.ctx.Err() != nil {
			workflowRun.Status = domain.WorkflowRunStatusTypeCanceled
		}
		workflowRun.EndedAt = time.Now()
		if invokeErr != nil {
			workflowRun.Error = embeddedhooks.SafeError(task.ctx, invokeErr).Error()
		} else {
			workflowRun.Error = "workflow did not reach a terminal state"
		}
		_, _ = wd.workflowRunRepo.SaveWithCascading(context.Background(), workflowRun)
	}

	wd.syslog.Info(fmt.Sprintf("workflow #%s's run #%s stopped", task.WorkflowId, task.RunId))
}

// tryNextLocked reserves each run atomically. Canceled tasks retain their
// workflow slot until their goroutine exits, so successors cannot overlap.
func (wd *workflowDispatcher) tryNextLocked() {
	if !wd.booted {
		return
	}
	for i := 0; i < len(wd.pendingRunQueue); {
		if wd.concurrency > 0 && len(wd.processingTasks) >= wd.concurrency {
			return
		}
		id := wd.pendingRunQueue[i]
		run, err := wd.workflowRunRepo.GetById(context.Background(), id)
		if err != nil || run.Status != domain.WorkflowRunStatusTypePending {
			wd.pendingRunQueue = append(wd.pendingRunQueue[:i], wd.pendingRunQueue[i+1:]...)
			continue
		}
		busy := false
		for _, task := range wd.processingTasks {
			if task.WorkflowId == run.WorkflowId {
				busy = true
				break
			}
		}
		if busy {
			i++
			continue
		}
		ctx, cancel := context.WithCancel(embeddedhooks.Bind(context.Background()))
		task := &taskInfo{WorkflowId: run.WorkflowId, RunId: run.Id, ctx: ctx, cancel: cancel}
		wd.pendingRunQueue = append(wd.pendingRunQueue[:i], wd.pendingRunQueue[i+1:]...)
		wd.processingTasks[id] = task
		wd.taskWG.Add(1)
		go wd.tryExecuteAsync(task)
	}
}

// New creates a dispatcher scoped to the currently bound application.
func New() WorkflowDispatcher { return newWorkflowDispatcher() }

func newWorkflowDispatcher() WorkflowDispatcher {
	return &workflowDispatcher{
		concurrency: envMaxWorkers,

		pendingRunQueue: make([]string, 0),
		processingTasks: make(map[string]*taskInfo),

		workflowRepo:    repository.NewWorkflowRepository(),
		workflowRunRepo: repository.NewWorkflowRunRepository(),
		workflowLogRepo: repository.NewWorkflowLogRepository(),

		syslog: app.GetLogger(),
	}
}
