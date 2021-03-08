package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/phungvandat/source-template/utils/errs"
)

// EncodeHTTPFunc define type for encode http
type EncodeHTTPFunc func(ctx context.Context, res http.ResponseWriter, resData interface{}, err error) error

// DecodeHTTPFunc define type for decode http
type DecodeHTTPFunc func(context.Context, *http.Request) (interface{}, error)

// StatusCoder interface
type StatusCoder interface {
	StatusCode() int
}

// EncodeJSONResponse function common with EncodeHTTPFunc type
func EncodeJSONResponse(ctx context.Context, resW http.ResponseWriter, res interface{}, err error) error {
	resW.Header().Set("Content-Type", "application/json; charset=utf-8")
	code := http.StatusOK

	if err != nil {
		var (
			httpCode = http.StatusInternalServerError
		)
		if sc, ok := err.(errs.CustomErr); ok {
			customCode = sc.StatusCode()
			if sc, ok := err.(StatusCoder); ok {
				code = sc.StatusCode()
			}

			if httpCode >= http.StatusContinue &&
				httpCode <= http.StatusNetworkAuthenticationRequired {
				httpCode = customCode
			}
		}
		resW.WriteHeader(code)
		return json.NewEncoder(resW).Encode(map[string]interface{}{
			"error": map[string]interface{}{
				"message": err.Error(),
				"code":    customCode,
			},
		})
	}

	if sc, ok := res.(StatusCoder); ok {
		code = sc.StatusCode()
	}

	resW.WriteHeader(code)
	if code == http.StatusNoContent {
		return nil
	}
	return json.NewEncoder(resW).Encode(res)
}
