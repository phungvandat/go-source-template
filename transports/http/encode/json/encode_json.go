package encode

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/phungvandat/source-template/pkg/errpkg"
	"github.com/phungvandat/source-template/utils/ctxkey"
)

// swagger:model ErrorResponse
type errRes struct {
	// Message of error
	Message string `json:"message"`
	// Code of error
	Code int `json:"code,omitempty"`
	// Type of error
	Type string `json:"type,omitempty"`
	// ID of trace error
	TraceID string `json:"trace_id,omitempty"`
}

func EncodeJSONResponse(ctx context.Context, w http.ResponseWriter, res interface{}, err error) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	httpCode := http.StatusOK

	if err != nil {
		httpCode = http.StatusInternalServerError
		var (
			errCode        int
			errMessage     = err.Error()
			type_, traceID string
		)
		cErr, ok := err.(errpkg.CustomErrorer)
		if ok {
			var errStatusCode = cErr.HTTPCode()
			if errStatusCode > 0 {
				httpCode = errStatusCode
			}

			errCode = cErr.Code()
			lang := ctxkey.GetLang(ctx)
			if lang != "" {
				errMessage = cErr.GetMessageByLang(errpkg.Lang(lang))
			}

			traceID = cErr.TraceID()
			type_ = cErr.Key()
		}

		w.WriteHeader(httpCode)
		// encode json response
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": errRes{
				Message: errMessage,
				Code:    errCode,
				Type:    type_,
				TraceID: traceID,
			},
		})
		return nil
	}

	w.WriteHeader(httpCode)
	if httpCode == http.StatusNoContent {
		return nil
	}

	return json.NewEncoder(w).Encode(res)
}
