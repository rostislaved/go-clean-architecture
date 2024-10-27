package http_adapter

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/http-adapter/handlers"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/http-adapter/router"
	"github.com/rostislaved/go-clean-architecture/internal/app/application/usecases"
	"github.com/rostislaved/go-clean-architecture/internal/app/config"
	http_server "github.com/rostislaved/go-clean-architecture/internal/libs/http-server"
)

type HttpAdapter struct {
	server *http_server.Server
}

func New(logger *slog.Logger, config config.HttpAdapter, svc *usecases.UseCases) *HttpAdapter {
	router := newRouter(logger, config.Router, svc)

	s := http_server.New(logger, config.Server, router)

	return &HttpAdapter{
		server: s,
	}
}

func newRouter(logger *slog.Logger, config config.Router, svc *usecases.UseCases) http.Handler {
	r := router.New()

	ctr := handlers.New(logger, svc)

	r.AppendRoutes(config, ctr)

	router := r.Router()

	return router
}

func (a HttpAdapter) Start(ctx context.Context) error {
	return a.server.Start(ctx)
}
