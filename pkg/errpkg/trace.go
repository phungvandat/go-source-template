package errpkg

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/phungvandat/source-template/utils/helper"
)

// ErrTracer to trace error and centralize error
type ErrTracer interface {
	Trace(err error, logErrs ...error) error
	GotErr() <-chan error
	Close()
	private()
}

type errTrace struct {
	errChn chan error
}

// NewErrTracer is constructor of ErrTracer
func NewErrTracer(maxChan int) ErrTracer {
	errChn := make(chan error, maxChan)

	return &errTrace{
		errChn: errChn,
	}
}

func (et *errTrace) private() {
	// Anti tampering
}

func (et *errTrace) Trace(err error, logErrs ...error) error {
	var (
		bufNum   = 3
		traceMsg = getStackTrace(bufNum)
		tErr, ok = err.(CustomErrorer)
	)

	if !ok {
		tErr = NewCustomErrByMsg(err.Error(), Option{HTTPCode: http.StatusInternalServerError})
	}

	if tErr.TraceID() == "" {
		tErr.SetTraceID()
		go helper.Goroutine(func() {
			if et.errChn == nil {
				return
			}
			logErrMsg := ""
			for idx := range logErrs {
				if logErrs[idx] != nil {
					logErrMsg += fmt.Sprintf("\nsecondary error: %v", logErrs[idx].Error())
				}
			}
			nErr := fmt.Errorf("%v \ntrace_id: %v%v%v", tErr, tErr.TraceID(), traceMsg, logErrMsg)
			et.errChn <- nErr
		})
	}

	return tErr
}

func (et *errTrace) GotErr() <-chan error {
	return et.errChn
}

// Close tracer
func (et *errTrace) Close() {
	close(et.errChn)
}

func getStackTrace(bufNum int) string {
	var (
		stackBuf = make([]uintptr, bufNum)
		skip     = 3
		length   = runtime.Callers(skip, stackBuf)
		stack    = stackBuf[:length]
		trace    = ""
		frames   = runtime.CallersFrames(stack)
	)

	for {
		frame, more := frames.Next()
		trace += fmt.Sprintf("\n%s:%d, Function: %s. ", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}

	return trace
}
