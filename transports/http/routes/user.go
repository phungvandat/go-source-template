package routes

import (
	"github.com/go-chi/chi"
	httpkit "github.com/go-kit/kit/transport/http"
	"github.com/phungvandat/source-template/endpoints"
	"github.com/phungvandat/source-template/transports/http/encode"
)

// UserRoutes define all route relate to user
func UserRoutes(
	endpoints endpoints.Endpoints,
	options []httpkit.ServerOption,
) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", httpkit.NewServer(
			endpoints.User.Login,
			nil,
			encode.JSONResponse,
			options...,
		).ServeHTTP)
	}
}
