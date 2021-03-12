package http

import (
	"context"
	"net/http"

	"github.com/phungvandat/source-template/pkg/endpoint"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request) error

func NewHandle(
	decode DecodeFunc,
	endpointHandler endpoint.Endpoint,
	encode EndcodeFunc,
) HandleFunc {
	ctx := context.Background()
	return func(w http.ResponseWriter, r *http.Request) error {
		in, err := decode(ctx, r)
		if err != nil {
			return err
		}

		out, err := endpointHandler(ctx, in)
		if err != nil {
			return err
		}

		return encode(ctx, w, out, err)
	}
}
