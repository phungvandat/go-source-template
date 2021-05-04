package helper

import (
	"fmt"
	"reflect"
	"runtime/debug"

	"github.com/phungvandat/source-template/utils/logger"
)

func Goroutine(method interface{}, args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			msg := fmt.Sprintf("%v\n%v", r, string(debug.Stack()))
			logger.Panic(msg)
		}
	}()

	if reflect.TypeOf(method).Kind() != reflect.Func {
		logger.Error("Method must function")
		return
	}

	vArgs := make([]reflect.Value, len(args))
	for i, val := range args {
		vArgs[i] = reflect.ValueOf(val)
	}

	reflect.ValueOf(method).Call(vArgs)
}
