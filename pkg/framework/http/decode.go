package http

import (
	"context"
	"net/http"
)

type DecodeFunc func(context.Context, *http.Request) (interface{}, error)
