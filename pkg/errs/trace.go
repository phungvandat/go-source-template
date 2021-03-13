package errs

import (
	"fmt"
	"net/http"
	"runtime"
)

// ErrTracer to trace error and centralize error
type ErrTracer interface {
	Trace(err error) error
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

func (et *errTrace) Trace(err error) error {
	var (
		bufNum   = 5
		traceMsg = getStackTrace(bufNum)
		tErr, ok = err.(CustomErrorer)
	)

	if !ok {
		tErr = NewCustomErrByMsg(err.Error(), Option{HTTPStatusCode: http.StatusInternalServerError})
	}

	if !tErr.IsTraced() {
		go func(err error) {
			if et.errChn == nil {
				return
			}
			nErr := fmt.Errorf("%v \nTrace: %v", err.Error(), traceMsg)
			et.errChn <- nErr
		}(tErr)
	}

	tErr.SetTraced(true)
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
		trace += fmt.Sprintf("\n\tFile: %s, Function: %s, Line: %d. ", frame.File, frame.Function, frame.Line)
		if !more {
			break
		}
	}

	return trace
}
