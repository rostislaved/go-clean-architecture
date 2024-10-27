package pprof_adapter

import (
	"context"
	"log/slog"
	"net/http/pprof"

	"github.com/gorilla/mux"

	http_server "github.com/rostislaved/go-clean-architecture/internal/libs/http-server"
)

type PprofAdapter struct {
	server *http_server.Server
}

type Config struct {
	Server http_server.Config
}

func New(logger *slog.Logger, config Config) *PprofAdapter {
	router := newPprofRouter()

	s := http_server.New(logger, config.Server, router)

	return &PprofAdapter{
		server: s,
	}
}

func newPprofRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)

	return router
}

func (a PprofAdapter) Start(ctx context.Context) error {
	return a.server.Start(ctx)
}
