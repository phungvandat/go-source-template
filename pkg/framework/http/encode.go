package http

import (
	"context"
	"net/http"
)

type EndcodeFunc func(ctx context.Context, res http.ResponseWriter, resData interface{}, err error) error
