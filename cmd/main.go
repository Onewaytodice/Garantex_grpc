package main

import (
	"Garantex_grpc/internal/config"
	grpcserver "Garantex_grpc/internal/grpc_server"
	"Garantex_grpc/internal/service"
	"Garantex_grpc/internal/storage/postgres"
	"Garantex_grpc/internal/web_client/garantex"
	"Garantex_grpc/pkg/logger"
	pbexchange "Garantex_grpc/proto/exchange_v1"
	"context"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"

	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.MustInit(conf.Logger.Name, conf.Logger.Production)
	logger.Info("Configuration loaded")

	// Connect to postgreSQL db
	db, err := conf.Database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	logger.Info("Database connect successful")

	// Initialize OTLP tracing
	provider, err := conf.Trace.InitTracerProvider(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = provider.Shutdown(context.Background())
	}()
	logger.Info("Tracing provider initialized")

	// Initialize metrics
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		logger.Info("HTTP metrics server running",
			zap.String("host", conf.HTTPServer.Host),
			zap.String("port", conf.HTTPServer.Port))

		if err := http.ListenAndServe(config.CreateAddr(conf.HTTPServer.Host, conf.HTTPServer.Port), nil); err != nil {
			logger.Error("failed to serve metrics http server", zap.Error(err))
		}
	}()

	// Up migrations
	err = conf.Database.UpMigrations("file://./migrations")
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Migrations successful")

	// Initializing app components
	web := garantex.NewGarantex(logger, conf.Garantex)

	storage := postgres.NewStorage(logger, db)

	exchange := service.NewExchange(logger, web, storage)

	grpcExchange := grpcserver.NewExchange(logger, exchange)

	// Initialize gRPC server
	srv := grpc.NewServer()
	pbexchange.RegisterExchangeGRPCServer(srv, grpcExchange)

	// Register healthcheck on gRPC server
	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(srv, healthcheck)

	// Register reflection service on gRPC server
	reflection.Register(srv)

	l, err := net.Listen("tcp", config.CreateAddr(conf.GRPCServer.Host, conf.GRPCServer.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	go func() {
		logger.Info("gRPC server running",
			zap.String("host", conf.GRPCServer.Host),
			zap.String("port", conf.GRPCServer.Port))

		if err := srv.Serve(l); err != nil {
			logger.Error("failed to serve gRPC server", zap.Error(err))
		}
	}()

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	logger.Warn("Received shutdown signal", zap.String("signal", sig.String()))

	// gRPC server graceful shutdown
	srv.GracefulStop()

	logger.Info("Server stopped gracefully")
}
