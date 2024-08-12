package router

import (
	"net/http"

	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/http-adapter"
	"github.com/rostislaved/go-clean-architecture/internal/pkg/middleware-helpers"

	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/http-adapter/api-controller"
)

func (r *Router) AppendRoutes(config httpAdapter.RouterConfig, controller *apiController.Controller) {
	r.config = config

	apiV1Subrouter := r.router.PathPrefix(apiV1Prefix).Subrouter()

	routes := []Route{
		{
			Name:    "method1",
			Path:    "/method1",
			Method:  http.MethodPost,
			Handler: middlewarehelpers.And()(http.HandlerFunc(controller.Get)),
		},
	}

	r.appendRoutesToRouter(apiV1Subrouter, routes)
}
