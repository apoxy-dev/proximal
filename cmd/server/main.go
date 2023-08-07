package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/temporalio/temporalite"
	tclient "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	envoyrunner "github.com/apoxy-dev/proximal/core/envoy"
	"github.com/apoxy-dev/proximal/core/log"
	"github.com/apoxy-dev/proximal/core/server"
	"github.com/apoxy-dev/proximal/server/api"
	serverdb "github.com/apoxy-dev/proximal/server/db"
	"github.com/apoxy-dev/proximal/server/envoy"
	"github.com/apoxy-dev/proximal/server/ingest"
	"github.com/apoxy-dev/proximal/server/watcher"

	endpointv1 "github.com/apoxy-dev/proximal/api/endpoint/v1"
	logsv1 "github.com/apoxy-dev/proximal/api/logs/v1"
	middlewarev1 "github.com/apoxy-dev/proximal/api/middleware/v1"
)

var (
	temporalHost = flag.String("temporal_host", "localhost:8088", "Temporal host.")
	temporalNs   = flag.String("temporal_namespace", "default", "Temporal namespace.")

	localMode    = flag.Bool("local_mode", true, "Run in local mode. (Launches a Temporalite server.)")
	buildDir     = flag.String("build_work_dir", "/tmp/proximal", "Base path for build artifacts.")
	watchIgnores = flag.String("watch_ignores", "^README.md$", "Comma-separated list of paths to ignore when watching for changes.")

	envoyPath       = flag.String("envoy_path", "", "Path to Envoy binary.")
	xdsListenHost   = flag.String("xds_host", "0.0.0.0", "XDS host.")
	xdsListenPort   = flag.Int("xds_port", 18000, "XDS port.")
	xdsSyncInterval = flag.Duration("xds_sync_interval", 10*time.Second, "How often to sync XSDs snapshot.")
)

func main() {
	s := server.NewApoxyServer(
		server.WithHandlers(middlewarev1.RegisterMiddlewareServiceHandlerFromEndpoint),
		server.WithHandlers(logsv1.RegisterLogsServiceHandlerFromEndpoint),
		server.WithHandlers(endpointv1.RegisterEndpointServiceHandlerFromEndpoint),
	)

	var ts *temporalite.Server
	var err error
	if *localMode {
		if ts, err = createTemporaliteServer(); err != nil {
			log.Fatalf("error creating Temporalite Server: %v", err)
		}
		go func() {
			if err := ts.Start(); err != nil {
				log.Fatalf("error starting Temporalite Server: %v", err)
			}
		}()
	}

	tc, err := dialTemporalWithRetries(s.Context, *temporalHost, *temporalNs)
	if err != nil {
		log.Fatalf("tclient.Dial() error: %v", err)
	}
	defer tc.Close()

	wOpts := worker.Options{
		MaxConcurrentActivityExecutionSize:     runtime.NumCPU(),
		MaxConcurrentWorkflowTaskExecutionSize: runtime.NumCPU(),
		EnableSessionWorker:                    true,
	}
	w := worker.New(tc, ingest.MiddlewareIngestQueue, wOpts)

	w.RegisterWorkflow(ingest.StartBuildWorkflow)
	w.RegisterWorkflow(ingest.DoBuildWorkflow)

	db, err := serverdb.New()
	if err != nil {
		log.Fatalf("serverdb.New() error: %v", err)
	}
	defer db.Close()
	if err = serverdb.DoMigrations(); err != nil {
		log.Fatalf("serverdb.DoMigrations() error: %v", err)
	}

	grpcClient, err := grpc.DialContext(
		s.Context,
		fmt.Sprintf("127.0.0.1:%d", *server.GRPCPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	fw := watcher.NewWatcher(
		middlewarev1.NewMiddlewareServiceClient(grpcClient),
		*watchIgnores,
	)
	go func() {
		if err := fw.Run(s.Context); err != nil {
			log.Errorf("fw.Run() error: %v", err)
		}
	}()

	ww := ingest.NewIngestWorker(*buildDir, db, fw)
	w.RegisterActivity(ww.PrepareGithubBuildActivity)
	w.RegisterActivity(ww.PrepareLocalBuildActivity)
	w.RegisterActivity(ww.BuildActivity)
	w.RegisterActivity(ww.UploadWasmOutputActivity)
	w.RegisterActivity(ww.FinalizeActivity)
	go func() {
		err = w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalf("w.Run() error: %v", err)
		}
	}()

	lsvc := api.NewLogsService(db)
	lsvc.RegisterALS(s.GRPC)

	middlewarev1.RegisterMiddlewareServiceServer(s.GRPC, api.NewMiddlewareService(db, tc, fw))
	logsv1.RegisterLogsServiceServer(s.GRPC, lsvc)
	endpointv1.RegisterEndpointServiceServer(s.GRPC, api.NewEndpointService(db, tc))

	envoyMgr := envoy.NewSnapshotManager(
		s.Context,
		middlewarev1.NewMiddlewareServiceClient(grpcClient),
		endpointv1.NewEndpointServiceClient(grpcClient),
		*buildDir,
		*xdsListenHost,
		*xdsListenPort,
		*xdsSyncInterval,
	)
	envoyMgr.RegisterXDS(s.GRPC)
	go func() {
		if err := envoyMgr.Run(s.Context); err != nil {
			log.Fatalf("envoyMgr.Run() error: %v", err)
		}
	}()

	e := &envoyrunner.Runtime{
		EnvoyPath:           *envoyPath,
		BootstrapConfigPath: "./cmd/server/envoy-bootstrap.yaml",
		Release: &envoyrunner.Release{
			Version: "v1.26.3",
		},
	}
	go func() {
		if err := e.Run(s.Context); err != nil {
			log.Fatalf("e.Run() error: %v", err)
		}
	}()

	go func() {
		exitCh := make(chan os.Signal, 1) // Buffered because sender is not waiting.
		signal.Notify(exitCh, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-exitCh:
		case <-s.Context.Done():
		}

		if *localMode {
			ts.Stop()
		}

		s.Shutdown()
	}()

	s.Run()
}

func dialTemporalWithRetries(ctx context.Context, hostPort, namespace string) (tclient.Client, error) {
	var tc tclient.Client
	var err error
	for i := 0; i < 10; i++ {
		tc, err = tclient.Dial(tclient.Options{
			HostPort:  hostPort,
			Namespace: namespace,
			ConnectionOptions: tclient.ConnectionOptions{
				EnableKeepAliveCheck: true,
			},
		})
		if err == nil {
			return tc, nil
		}
		log.Errorf("tclient.Dial() error: %v", err)
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(time.Second):
		}
	}
	return nil, err
}
