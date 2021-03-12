package mws

import (
	"net/http"
	"time"

	"github.com/phungvandat/source-template/utils/logger"
	"github.com/phungvandat/source-template/utils/timeutil"
)

func RequestInfo(r *http.Request, startTime time.Time, errMsg *string) {
	template := "Method: %v, Route: %v, Time: %v Duration: %v ms"
	endMillis := timeutil.GetNowMillis()
	args := []interface{}{
		r.Method,
		r.RequestURI,
		startTime.Format("2006/01/02-15:04:05"),
		endMillis - timeutil.TimeToMillis(startTime),
	}

	if errMsg != nil && *errMsg != "" {
		template += " Err: %v"
		args = append(args, *errMsg)
		logger.Error(template, args...)
		return
	}
	logger.Info(template, args...)
}
