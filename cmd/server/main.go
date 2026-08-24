// Command server wires together config, the Postgres pool, the cedar-agent
// sidecar client, the in-memory policy cache, the async audit logger, and
// starts both the REST and gRPC listeners.
package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/kwasi/policy-engine/internal/audit"
	"github.com/kwasi/policy-engine/internal/cache"
	"github.com/kwasi/policy-engine/internal/cedarclient"
	"github.com/kwasi/policy-engine/internal/config"
	grpcapi "github.com/kwasi/policy-engine/internal/api/grpc"
	restapi "github.com/kwasi/policy-engine/internal/api/rest"
	"github.com/kwasi/policy-engine/internal/storage"
	"github.com/kwasi/policy-engine/internal/telemetry"
)

func main() {
	cfg := config.Load()
	logger := telemetry.NewLogger(cfg.LogLevel)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	pool, err := storage.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		logger.Error("failed to connect to postgres", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	cedarClient := cedarclient.New(cfg.CedarAgentURL, cfg.CedarAgentToken)
	policyRepo := storage.NewPolicyRepo(pool)
	auditRepo := storage.NewAuditRepo(pool)
	policyCache := cache.New(policyRepo, cedarClient, cfg.PolicyCacheTTL)
	auditor := audit.New(auditRepo, 1024)

	if _, err := policyCache.Invalidate(ctx); err != nil {
		logger.Error("initial policy sync to cedar-agent failed", "error", err)
		os.Exit(1)
	}

	restServer := restapi.NewServer(cedarClient, policyRepo, policyCache, auditor)
	grpcServer := grpc.NewServer()
	grpcapi.NewServer(cedarClient, policyRepo, policyCache, auditor).Register(grpcServer)

	go func() {
		logger.Info("rest listening", "addr", cfg.RESTListenAddr)
		if err := http.ListenAndServe(cfg.RESTListenAddr, restServer.Routes()); err != nil {
			logger.Error("rest server stopped", "error", err)
		}
	}()

	go func() {
		lis, err := net.Listen("tcp", cfg.GRPCListenAddr)
		if err != nil {
			logger.Error("grpc listen failed", "error", err)
			return
		}
		logger.Info("grpc listening", "addr", cfg.GRPCListenAddr)
		if err := grpcServer.Serve(lis); err != nil {
			logger.Error("grpc server stopped", "error", err)
		}
	}()

	<-ctx.Done()
	logger.Info("shutting down")
	grpcServer.GracefulStop()
}
