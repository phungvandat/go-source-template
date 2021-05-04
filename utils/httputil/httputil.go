package httputil

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"

	"github.com/phungvandat/source-template/utils/errs"
	"github.com/phungvandat/source-template/utils/logger"
)

func DecodeJSON(r *http.Request, out interface{}) error {
	if out == nil || reflect.TypeOf(out).Kind() != reflect.Ptr {
		logger.Error("invalid type to decode json")
		return errs.ErrSomethingWentWrong
	}

	err := json.NewDecoder(r.Body).Decode(out)
	if err != nil && err != io.EOF {
		logger.Error("failed decode json by error: %v", err)
		return errs.ErrBodyNotAllowed
	}

	return nil
}
