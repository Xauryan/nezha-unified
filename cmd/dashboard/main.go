package main

import (
	"context"
	"crypto/tls"
	"embed"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/debug"
	"strings"
	"syscall"
	"time"
	_ "time/tzdata"

	"github.com/certimate-go/certimate/pkg/embedded"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/nezhahq/nezha/cmd/dashboard/controller"
	"github.com/nezhahq/nezha/cmd/dashboard/controller/waf"
	"github.com/nezhahq/nezha/cmd/dashboard/rpc"
	"github.com/nezhahq/nezha/model"
	"github.com/nezhahq/nezha/pkg/idcodec"
	"github.com/nezhahq/nezha/pkg/utils"
	"github.com/nezhahq/nezha/proto"
	rpcService "github.com/nezhahq/nezha/service/rpc"
	"github.com/nezhahq/nezha/service/singleton"
	"github.com/nezhahq/nezha/service/unifiedbackup"
)

type DashboardCliParam struct {
	Version           bool
	ConfigFile        string
	DatabaseLocation  string
	StorageOnly       bool
	BackupCreate      string
	BackupRestore     string
	BackupDestination string
	BackupKeyFile     string
	BackupRehearse    string
	BackupInspect     string
	BackupReconcile   string
	BackupEvidence    string
}

var (
	dashboardCliParam DashboardCliParam
	// The process owns this lock until exit, including any legacy goroutine that
	// outlives HTTP shutdown. Offline maintenance must never overlap those writes.
	dashboardDataUnlock func() error
	//go:embed *-dist
	frontendDist embed.FS
)

func initSystem(bus chan<- *model.Service) error {
	var usersCount int64
	if err := singleton.DB.Model(&model.User{}).Count(&usersCount).Error; err != nil {
		return err
	}
	// Backward-compatible bootstrap state: existing installers and recovery
	// procedures expect the first login on an empty database to be admin/admin.
	// This is not a permanent credential or an authentication-bypass fallback;
	// operators must complete initialization and change it before exposing the
	// Dashboard. Replacing it requires a coordinated installer/migration flow so
	// existing unattended installations are not locked out.
	if usersCount == 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		admin := model.User{
			Username: "admin",
			Password: string(hash),
		}
		if err := singleton.DB.Create(&admin).Error; err != nil {
			return err
		}
	}

	if err := singleton.LoadSingleton(bus); err != nil {
		return err
	}

	if _, err := singleton.CronShared.AddFunc("0 30 3 * * *", singleton.CleanMonitorHistory); err != nil {
		return err
	}

	if _, err := singleton.CronShared.AddFunc("0 0 * * * *", func() { singleton.RecordTransferHourlyUsage() }); err != nil {
		return err
	}

	if err := singleton.StartJWTSessionGC(); err != nil {
		return err
	}
	return nil
}

func initIDCodec() error {
	return idcodec.Init([]byte(singleton.Conf.JWTSecretKey))
}

// @title           Nezha Monitoring API
// @version         1.0
// @description     Nezha Monitoring API
// @termsOfService  http://nezhahq.github.io

// @contact.name   API Support
// @contact.url    http://nezhahq.github.io
// @contact.email  hi@nai.ba

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8008
// @BasePath  /api/v1

// @securityDefinitions.apikey  BearerAuth
// @in header
// @name Authorization
// @description JWT session token. Browser/UI flow. Format: `Bearer <jwt>` or cookie `nz-jwt`.

// @securityDefinitions.apikey  APITokenAuth
// @in header
// @name Authorization
// @description Personal Access Token (PAT). Programmatic/CI/LLM flow. Format: `Bearer nzp_<secret>`.
// @description Each endpoint enforces a specific scope; see the `controller` package godoc for the authoritative scope table.

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	flag.BoolVar(&dashboardCliParam.Version, "v", false, "查看当前版本号")
	flag.StringVar(&dashboardCliParam.ConfigFile, "c", "data/config.yaml", "配置文件路径")
	flag.StringVar(&dashboardCliParam.DatabaseLocation, "db", "data/sqlite.db", "Sqlite3数据库文件路径")
	flag.BoolVar(&dashboardCliParam.StorageOnly, "storage-only", false, "Development only: disable certificate runtime and DDNS workers")
	registerBackupFlags()
	flag.Parse()

	if dashboardCliParam.Version {
		fmt.Println(singleton.Version)
		return
	}
	if handled, err := runBackupCommand(context.Background(), dashboardCliParam); handled {
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if err := runDashboard(); err != nil {
		log.Fatal(err)
	}
}

func runDashboard() error {
	if !dashboardCliParam.StorageOnly && !embedded.RuntimeAvailable() {
		return errors.New("this Dashboard was built without certificate runtime; use the full release build or explicit --storage-only for development")
	}
	signalContext, stopSignals := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stopSignals()
	dataRoot := filepath.Dir(dashboardCliParam.DatabaseLocation)
	unlock, err := unifiedbackup.LockDataRoot(dataRoot)
	if err != nil {
		return err
	}
	dashboardDataUnlock = unlock
	if info, err := os.Lstat(dashboardCliParam.DatabaseLocation); err == nil {
		if !info.Mode().IsRegular() {
			return errors.New("Dashboard database must be a regular file inside the locked data directory")
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}
	// Fence restored stores before migrations, schedulers or a mutating API.
	if err := unifiedbackup.CheckReady(dataRoot); err != nil {
		return err
	}
	lifetime, cancelLifetime := context.WithCancel(signalContext)
	defer cancelLifetime()
	shutdownAttempted := false
	serviceSentinelDispatchBus := make(chan *model.Service)
	if err := utils.FirstError(singleton.InitFrontendTemplates,
		func() error { return singleton.InitConfigFromPath(dashboardCliParam.ConfigFile) },
		initIDCodec,
		singleton.InitTimezoneAndCache,
		func() error {
			if singleton.Conf.Memory.GoMemLimitMB > 0 {
				debug.SetMemoryLimit(singleton.Conf.Memory.GoMemLimitMB * 1024 * 1024)
				log.Printf("NEZHA>> Go memory limit set to %d MB", singleton.Conf.Memory.GoMemLimitMB)
			}
			return nil
		},
		func() error { return singleton.InitDBFromPath(dashboardCliParam.DatabaseLocation) },
		singleton.InitTSDB,
		func() error { return initSystem(serviceSentinelDispatchBus) }); err != nil {
		return err
	}

	if err := singleton.InitUnified(filepath.Join(filepath.Dir(dashboardCliParam.DatabaseLocation), "certimate")); err != nil {
		return err
	}
	rpcService.SetUnifiedHostObserver(singleton.Unified.NotifyDDNSHost)
	rpcService.SetUnifiedHostIPObserver(singleton.ReportUnifiedHostIP)
	rpcService.SetUnifiedHostAddressesObserver(singleton.ReportUnifiedHostAddresses)
	defer rpcService.SetUnifiedHostObserver(nil)
	defer rpcService.SetUnifiedHostIPObserver(nil)
	defer rpcService.SetUnifiedHostAddressesObserver(nil)
	defer func() {
		cancelLifetime()
		if shutdownAttempted {
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := stopUnified(ctx); err != nil {
			log.Printf("NEZHA>> unified shutdown did not drain: %v", err)
		}
	}()
	if !dashboardCliParam.StorageOnly {
		if err := singleton.Unified.StartRuntime(lifetime); err != nil {
			return err
		}
		if err := singleton.Unified.StartDDNS(lifetime, singleton.UnifiedHostSource); err != nil {
			return err
		}
		singleton.Unified.SetAgentCalls(rpcService.UnifiedAgentCalls)
	} else {
		log.Println("NEZHA>> Development storage-only mode: certificate runtime and DDNS are disabled")
	}

	l, err := openDashboardListener("tcp", dashboardListenerAddress(singleton.Conf.ListenHost, singleton.Conf.ListenPort), dashboardHTTPListener)
	if err != nil {
		return err
	}
	defer l.Close()
	receiptListener, err := openReceiptGateListener()
	if err != nil {
		return err
	}
	if receiptListener != nil {
		defer receiptListener.Close()
		rpc.SetReceiptGateListener(receiptListener)
	}

	singleton.CleanMonitorHistory()
	rpc.DispatchKeepalive()
	rpc.SetMCPKillSwitchObserver(func() bool {
		return singleton.Conf == nil || !singleton.Conf.MCPEnabled()
	})
	go rpc.DispatchTask(serviceSentinelDispatchBus)
	go singleton.AlertSentinelStart()

	grpcHandler := rpc.ServeRPC()
	httpHandler := controller.ServeWeb(frontendDist)
	controller.InitUpgrader()
	wireNATDashboardCredentialGate()

	muxHandler := newHTTPandGRPCMux(httpHandler, grpcHandler)
	muxServerHTTP := &http.Server{
		Handler:           muxHandler,
		ReadHeaderTimeout: time.Second * 5,
	}
	muxServerHTTP.Protocols = new(http.Protocols)
	muxServerHTTP.Protocols.SetHTTP1(true)
	muxServerHTTP.Protocols.SetUnencryptedHTTP2(true)

	var muxServerHTTPS *http.Server
	if singleton.Conf.HTTPS.ListenPort != 0 {
		muxServerHTTPS = &http.Server{
			Addr:              fmt.Sprintf("%s:%d", singleton.Conf.ListenHost, singleton.Conf.HTTPS.ListenPort),
			Handler:           muxHandler,
			ReadHeaderTimeout: time.Second * 5,
			TLSConfig: &tls.Config{
				InsecureSkipVerify: singleton.Conf.HTTPS.InsecureTLS,
			},
		}
	}

	errChan := make(chan error, 2)
	log.Printf("NEZHA>> Dashboard::START ON %s:%d", singleton.Conf.ListenHost, singleton.Conf.ListenPort)
	if muxServerHTTPS != nil {
		go func() {
			errChan <- serveDashboardHTTPS(muxServerHTTPS, singleton.Conf.HTTPS.TLSCertPath, singleton.Conf.HTTPS.TLSKeyPath)
		}()
	}
	go func() { errChan <- muxServerHTTP.Serve(l) }()
	var serveErr error
	select {
	case <-signalContext.Done():
	case serveErr = <-errChan:
	}
	// Cancel schedulers/providers, terminate persistent RPC streams, then drain
	// HTTP before storage. Every shutdown step shares one deadline.
	cancelLifetime()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	log.Println("NEZHA>> Graceful::START")
	rpc.CloseReceiptGate()
	grpcHandler.Stop()
	var shutdownErr error
	if muxServerHTTPS != nil {
		shutdownErr = muxServerHTTPS.Shutdown(ctx)
	}
	shutdownErr = errors.Join(shutdownErr, muxServerHTTP.Shutdown(ctx))
	if shutdownErr != nil {
		_ = muxServerHTTP.Close()
		if muxServerHTTPS != nil {
			_ = muxServerHTTPS.Close()
		}
	}
	shutdownAttempted = true
	shutdownErr = errors.Join(shutdownErr, stopUnified(ctx))
	singleton.RecordTransferHourlyUsage()
	singleton.CloseTSDB()
	log.Println("NEZHA>> Graceful::END")
	if errors.Is(serveErr, http.ErrServerClosed) {
		serveErr = nil
	}
	return errors.Join(serveErr, shutdownErr)
}

// stopUnified leaves storage open after an incomplete drain so late operations
// retain their authorization hooks and can persist their terminal state.
func stopUnified(ctx context.Context) error {
	if singleton.Unified == nil {
		return nil
	}
	if err := singleton.Unified.StopDDNSContext(ctx); err != nil {
		return err
	}
	if err := singleton.Unified.CloseDeploymentsContext(ctx); err != nil {
		return err
	}
	if err := singleton.Unified.StopRuntime(ctx); err != nil {
		return err
	}
	return singleton.Unified.Close()
}

// wireNATDashboardCredentialGate connects the NAT ingress to the real
// dashboard-credential classifiers. The NAT ingress strips dashboard-issued
// credentials (panel JWT / panel PAT) from every channel the panel's own
// TokenLookup accepts — the Authorization header, the ?token= query parameter
// and the nz-jwt cookie — before tunneling a request to the agent, and
// forwards foreign values untouched. ServeWeb built the JWT parser the
// classifier reuses, so the gate is wired right after it and before any
// listener starts serving.
func wireNATDashboardCredentialGate() {
	rpc.SetNATDashboardCredentialGate(controller.ClassifyDashboardCredentialValues)
}

func newHTTPandGRPCMux(httpHandler http.Handler, grpcHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		natConfig := singleton.NATShared.GetNATConfigByDomain(r.Host)
		if natConfig != nil {
			if !natConfig.Enabled {
				c, _ := gin.CreateTestContext(w)
				waf.ShowBlockPage(c, fmt.Errorf("nat host %s is disabled", natConfig.Domain))
				return
			}
			rpc.ServeNAT(w, r, natConfig)
			return
		}
		if r.ProtoMajor == 2 && r.Header.Get("Content-Type") == "application/grpc" &&
			strings.HasPrefix(r.URL.Path, "/"+proto.NezhaService_ServiceDesc.ServiceName) {
			grpcHandler.ServeHTTP(w, r)
			return
		}
		httpHandler.ServeHTTP(w, r)
	})
}
