package encode

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/phungvandat/source-template/model/domain"
	"github.com/phungvandat/source-template/utils/ctxkey"
	"github.com/phungvandat/source-template/utils/errs"
)

// swagger:model ErrorResponse
type errRes struct {
	// Message of error
	Message string `json:"message"`
	// Code of error
	Code int `json:"code"`
}

type StatusCoder interface {
	HTTPStatusCode() int
}

func EncodeJSONResponse(ctx context.Context, w http.ResponseWriter, res interface{}, err error) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	httpCode := http.StatusOK

	if sc, ok := res.(StatusCoder); ok {
		httpCode = sc.HTTPStatusCode()
	}
	if err != nil {
		var (
			errCode    int
			errMessage = err.Error()
		)
		cErr, ok := err.(errs.CustomErrorer)
		if ok {
			errCode = cErr.Code()
			lang := ctxkey.GetStrValue(ctx, domain.CtxKeyLang)
			if lang != "" {
				errMessage = cErr.GetMessageByLang(errs.ErrLang(lang))
			}
		}

		w.WriteHeader(httpCode)
		// enforce json response
		json.NewEncoder(w).Encode(errRes{
			Message: errMessage,
			Code:    errCode,
		})
		return nil
	}

	w.WriteHeader(httpCode)
	if httpCode == http.StatusNoContent {
		return nil
	}

	return json.NewEncoder(w).Encode(res)
}
