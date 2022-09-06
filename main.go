package main

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/iamgafurov/journal/internal/config"
	"github.com/iamgafurov/journal/internal/handler/rest"
	"github.com/iamgafurov/journal/internal/logger"
	"github.com/iamgafurov/journal/internal/service"
	"github.com/iamgafurov/journal/internal/storage/postgres"
	"github.com/iamgafurov/journal/internal/storage/sqlserver"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger.InitLogger()

	cfg, err := config.New()
	if err != nil {
		logger.Logger.Fatal("Config parsing", zap.Error(err))
	}

	//init sentry for handling internal errors
	err = sentry.Init(sentry.ClientOptions{
		Dsn: cfg.SentryDSN,
	})

	ctx, cancel := context.WithCancel(context.Background())

	postgresDB, err := postgres.New(cfg.PostgresConnStr)
	if err != nil {
		logger.Logger.Fatal("PostgresDB initialization", zap.Error(err))
	}

	mssqlDB, err := sqlserver.New(ctx, cfg.MSSQLConnStr)
	if err != nil {
		logger.Logger.Fatal("MSSQLDB initialization", zap.Error(err))
	}

	svc := service.New(postgresDB, mssqlDB, cfg, logger.Logger)

	server := rest.New(cfg, svc)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		logger.Logger.Info("Server shutting down")
		server.Shutdown(ctx)
		signal.Stop(sig)
		close(sig)
		cancel()
	}()
	server.Run()
}
