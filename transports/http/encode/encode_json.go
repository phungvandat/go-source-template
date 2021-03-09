package encode

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/phungvandat/source-template/utils/contextkey"
	"github.com/phungvandat/source-template/utils/errs"
)

// JSONResponse func
func JSONResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}

// JSONError func
func JSONError(ctx context.Context, err error, w http.ResponseWriter) {
	encodeJSONError(ctx, err, w)
}

func encodeJSONError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// custom headers
	if headerer, ok := err.(kithttp.Headerer); ok {
		for k, values := range headerer.Headers() {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
	}
	var (
		httpCode   = http.StatusInternalServerError
		errCode    int
		errMessage = err.Error()
	)
	cErr, ok := err.(errs.CustomErrorer)
	if ok {
		errCode = cErr.Code()
		if cErr.HTTPStatusCode() > 0 {
			httpCode = cErr.HTTPStatusCode()
		}
		lang := contextkey.GetLang(ctx)
		errMessage = cErr.GetMessageByLang(errs.ErrLang(lang))
	}

	w.WriteHeader(httpCode)
	// enforce json response
	json.NewEncoder(w).Encode(errRes{
		Message: errMessage,
		Code:    errCode,
	})
}

// swagger:model ErrorResponse
type errRes struct {
	// Message of error
	Message string `json:"message"`
	// Code of error
	Code int `json:"code"`
}
