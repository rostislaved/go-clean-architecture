package router

import (
	"net/http"

	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/http-adapter/handlers"
	middlewarehelpers "github.com/rostislaved/go-clean-architecture/internal/libs/middleware-helpers"
)

func (r *Router) AppendRoutes(config Config, handlers *handlers.Handlers) {
	r.config = config

	apiV1Subrouter := r.router.PathPrefix(apiV1Prefix).Subrouter()

	routes := []Route{
		{
			Name:    "method1",
			Path:    "/method1",
			Method:  http.MethodPost,
			Handler: middlewarehelpers.And()(http.HandlerFunc(handlers.Get)),
		},
	}

	r.appendRoutesToRouter(apiV1Subrouter, routes)
}
