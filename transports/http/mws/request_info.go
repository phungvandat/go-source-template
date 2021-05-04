package mws

import (
	"net/http"
	"time"

	"github.com/phungvandat/source-template/utils/helper"
	"github.com/phungvandat/source-template/utils/logger"
)

func RequestInfo(r *http.Request, startTime time.Time, errMsg *string) {
	template := "Method: %v, Route: %v, Time: %v Duration: %v ms"
	endMillis := helper.GetNowMillis()
	args := []interface{}{
		r.Method,
		r.RequestURI,
		startTime.Format("2006/01/02-15:04:05"),
		endMillis - helper.TimeToMillis(startTime),
	}

	if errMsg != nil && *errMsg != "" {
		template += " Err: %v"
		args = append(args, *errMsg)
		logger.Error(template, args...)
		return
	}
	logger.Info(template, args...)
}
