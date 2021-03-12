package http

import (
	"github.com/phungvandat/source-template/endpoints"
	fwk "github.com/phungvandat/source-template/pkg/framework/http"
	"github.com/phungvandat/source-template/transports/http/routes"
)

func BuildRouter(eps endpoints.Endpoints) fwk.Router {
	r := fwk.NewRouter()
	r.Group("/api", func(r fwk.Router) {
		r.Group("/internal", func(r fwk.Router) {
			r.Group("/v1", func(r fwk.Router) {
				r.Group("/users", routes.UserRoutes(eps))
			})
		})
	})
	return r
}
