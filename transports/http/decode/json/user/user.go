package user

import (
	"context"
	"net/http"

	"github.com/phungvandat/source-template/utils/httputil"

	iom "github.com/phungvandat/source-template/model/service/user"
)

// LoginRequest decode json data
func LoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		body = &iom.LoginIn{}
		err  = httputil.DecodeJSON(r, body)
	)

	return body, err
}
