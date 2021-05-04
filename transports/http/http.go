package http

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	fwk "github.com/phungvandat/source-template/pkg/framework/http"
	"github.com/phungvandat/source-template/transports/http/mws"
	"github.com/phungvandat/source-template/utils/errs"
	"github.com/phungvandat/source-template/utils/logger"
)

type httpHandler struct {
	mapRouteHandle map[string]fwk.HandleFunc
}

func NewHTTPHandler(mrh map[string]fwk.HandleFunc) http.Handler {
	return &httpHandler{
		mapRouteHandle: mrh,
	}
}

func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Panic Recovery
	defer func() {
		if r := recover(); r != nil {
			msg := fmt.Sprintf("%v\n%v", r, string(debug.Stack()))
			logger.Panic(msg)
		}
	}()

	// TODO: add cors when run in production
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Language")

	var (
		startTime = time.Now()
		errMsg    = new(string)
	)

	defer mws.RequestInfo(r, startTime, errMsg)
	defer func() {
		err := r.Body.Close()
		if err != nil {
			*errMsg = err.Error()
		}
	}()

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		*errMsg = errs.ErrBodyNotAllowed.Error()
		return
	}

	var (
		urlPath    = r.URL.Path
		handle, ok = h.mapRouteHandle[urlPath]
	)

	if !ok {
		*errMsg = errs.ErrRouteNotFound.Error()
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var err = handle(w, r)
	if err != nil {
		*errMsg = err.Error()
	}
}
