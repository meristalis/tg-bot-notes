// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/meristalis/tg-bot-notes/config"
	amqprpc "github.com/meristalis/tg-bot-notes/internal/controller/amqp_rpc"
	router "github.com/meristalis/tg-bot-notes/internal/controller/http"
	"github.com/meristalis/tg-bot-notes/internal/repo/persistent"
	"github.com/meristalis/tg-bot-notes/internal/repo/webapi"
	"github.com/meristalis/tg-bot-notes/internal/usecase/note"
	"github.com/meristalis/tg-bot-notes/internal/usecase/translation"
	"github.com/meristalis/tg-bot-notes/pkg/httpserver"
	"github.com/meristalis/tg-bot-notes/pkg/logger"
	"github.com/meristalis/tg-bot-notes/pkg/postgres"
	"github.com/meristalis/tg-bot-notes/pkg/rabbitmq/rmq_rpc/server"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	translationUseCase := translation.New(
		persistent.New(pg),
		webapi.New(),
	)

	NoteUseCase := note.New(
		persistent.NewNoteRepo(pg),
	)

	// RabbitMQ RPC Server
	rmqRouter := amqprpc.NewRouter(translationUseCase)

	rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	}

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port), httpserver.Prefork(cfg.HTTP.UsePreforkMode))
	router.NewRouter(httpServer.App, cfg, l, translationUseCase, NoteUseCase)

	// Start servers
	rmqServer.Start()
	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-rmqServer.Notify():
		l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	err = rmqServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	}
}
