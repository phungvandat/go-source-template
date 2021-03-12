package routes

import (
	"github.com/go-chi/chi"
	httpkit "github.com/go-kit/kit/transport/http"
	"github.com/phungvandat/source-template/endpoints"
	jsonDec "github.com/phungvandat/source-template/transports/http/decode/json/user"
	"github.com/phungvandat/source-template/transports/http/encode"
)

// UserRoutes define all route relate to user
func UserRoutes(
	endpoints endpoints.Endpoints,
	mws []httpkit.ServerOption,
) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", httpkit.NewServer(
			endpoints.User.Login,
			jsonDec.LoginRequest,
			encode.JSONResponse,
			mws...,
		).ServeHTTP)
	}
}
