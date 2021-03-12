package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/phungvandat/source-template/endpoints"
)

// NewHTTPHandler func create http handler
func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	var (
		r       = chi.NewRouter()
		origins = []string{"*"}
		headers = []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Language"}
	)

	r.Route("/api", func(r chi.Router) {
		r.Route("/internal", func(r chi.Router) {
			corsV1 := cors.New(cors.Options{
				AllowedOrigins:   origins,
				AllowedMethods:   []string{http.MethodPost, http.MethodOptions},
				AllowedHeaders:   headers,
				AllowCredentials: true,
			})
			r.Use(corsV1.Handler)
			r.Route("/v1", func(r chi.Router) {

			})
		})
	})
	return r
}
