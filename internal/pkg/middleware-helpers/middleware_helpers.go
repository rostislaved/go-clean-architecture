package middlewarehelpers

import "net/http"

type middleware = func(f http.Handler) http.Handler

func And(middlewares ...middleware) middleware {
	return func(h http.Handler) http.Handler {
		wrapped := h

		for _, middleware := range middlewares {
			wrapped = middleware(wrapped)
		}

		return wrapped
	}
}
