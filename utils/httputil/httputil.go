package httputil

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"

	"github.com/phungvandat/source-template/pkg/errs"
)

func DecodeJSON(r *http.Request, out interface{}) (interface{}, error) {
	if out == nil || reflect.TypeOf(out).Kind() != reflect.Ptr {
		return nil, errs.ErrInvalidTypeToDecodeJSON
	}

	err := json.NewDecoder(r.Body).Decode(out)
	if err != io.EOF {
		return out, nil
	}

	return out, nil
}
