// Package v1 implements routing paths. Each services in own file.
package http

import (
	"net/http"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/meristalis/tg-bot-notes/config"
	_ "github.com/meristalis/tg-bot-notes/docs" // Swagger docs.
	"github.com/meristalis/tg-bot-notes/internal/controller/http/middleware"
	v1 "github.com/meristalis/tg-bot-notes/internal/controller/http/v2"
	"github.com/meristalis/tg-bot-notes/internal/usecase"
	"github.com/meristalis/tg-bot-notes/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
func NewRouter(app *fiber.App, cfg *config.Config, l logger.Interface, t usecase.Translation, n usecase.Note) {
	// Options
	app.Use(middleware.Logger(l))
	app.Use(middleware.Recovery(l))

	// Prometheus metrics
	if cfg.Metrics.Enabled {
		prometheus := fiberprometheus.New("my-service-name")
		prometheus.RegisterAt(app, "/metrics")
		app.Use(prometheus.Middleware)
	}

	// Swagger
	if cfg.Swagger.Enabled {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	// K8s probe
	app.Get("/healthz", func(ctx *fiber.Ctx) error { return ctx.SendStatus(http.StatusOK) })

	// Routers
	app.Use(middleware.JWTMiddleware(cfg.Auth.PublicKey, l))

	apiV1Group := app.Group("/v1")
	{
		v1.NewNoteRoutes(apiV1Group, n, l)
	}
}
