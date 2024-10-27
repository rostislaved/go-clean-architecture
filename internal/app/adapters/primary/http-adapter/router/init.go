package router

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
	config Config
}

type Config struct {
	Shutdown             shutdown
	Timeout              timeout
	AuthenticationConfig string `config:"envVar"`
	AuthorizationConfig  string `config:"envVar"`
}

type shutdown struct {
	Duration time.Duration
}

type timeout struct {
	Duration time.Duration
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
