package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/config"
	middlewarehelpers "github.com/rostislaved/go-clean-architecture/internal/pkg/middleware-helpers"
)

type Router struct {
	router *mux.Router
	config config.Router
}

func New() *Router {
	router := mux.NewRouter()

	r := Router{
		router: router,
	}

	r.addInfrastructureRoutes()

	return &r
}

const (
	apiV1Prefix = "/api/v1"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.Handler
}

func (r *Router) Router() http.Handler {
	return r.router
}

func (r *Router) addInfrastructureRoutes() {
	r.router.
		Name("info").
		Methods(http.MethodGet).
		Path("/tech/info")

	r.router.
		Name("state").
		Methods(http.MethodGet).
		Path("/tech/state") // TODO

	r.router.
		Name("metrics").
		Methods(http.MethodGet).
		Path("/metrics").
		Handler(promhttp.Handler())
}

func (r *Router) appendRoutesToRouter(subrouter *mux.Router, routes []Route) {
	for i := range routes {
		routes[i].Handler = middlewarehelpers.And()(routes[i].Handler)
	}

	for _, route := range routes {
		subrouter.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Path).
			Handler(route.Handler)
	}
}
