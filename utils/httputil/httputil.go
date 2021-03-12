package httputil

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/phungvandat/source-template/utils/errs"
)

func DecodeJSON(r *http.Request, out interface{}) (interface{}, error) {
	if out == nil || reflect.TypeOf(out).Kind() != reflect.Ptr {
		return nil, errs.ErrInvalidTypeToDecodeJSON
	}

	err := json.NewDecoder(r.Body).Decode(out)

	return out, err
}
