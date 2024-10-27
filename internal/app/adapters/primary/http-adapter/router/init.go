package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rostislaved/go-clean-architecture/internal/app/config"
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

func (r *Router) appendRoutesToRouter(subrouter *mux.Router, routes []Route) {
	for _, route := range routes {
		subrouter.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Path).
			Handler(route.Handler)
	}
}
