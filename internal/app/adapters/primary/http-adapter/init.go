package httpAdapter

import (
	"log/slog"
	"net/http"
	"time"

	controller "github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/http-adapter/api-controller"
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/http-adapter/router"
	"github.com/rostislaved/go-clean-architecture/internal/app/application/service"
)

const (
	_defaultReadTimeout       = 5 * time.Second
	_defaultWriteTimeout      = 500 * time.Second
	_defaultReadHeaderTimeout = 5 * time.Second
	_defaultShutdownTimeout   = 3 * time.Second
)

type HttpAdapter struct {
	logger          *slog.Logger
	config          HttpAdapterConfig
	router          http.Handler
	server          *http.Server
	shutdownTimeout time.Duration
	notify          chan error
}

func New(logger *slog.Logger, config HttpAdapterConfig, svc *service.ApiService) HttpAdapter {
	r := router.New()

	ctr := controller.New(logger, svc)

	r.AppendRoutes(config.Router, ctr)

	router := r.Router()

	httpServer := &http.Server{
		Handler:           router,
		ReadTimeout:       _defaultReadTimeout,
		WriteTimeout:      _defaultWriteTimeout,
		ReadHeaderTimeout: _defaultReadHeaderTimeout,
		Addr:              config.Server.Port,
	}

	return HttpAdapter{
		logger:          logger,
		config:          config,
		router:          router,
		server:          httpServer,
		shutdownTimeout: _defaultShutdownTimeout,
		notify:          make(chan error, 1),
	}
}
