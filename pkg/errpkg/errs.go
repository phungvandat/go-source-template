package errpkg

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_1234567890")

// CustomErrorer interface
type CustomErrorer interface {
	Error() string
	HTTPCode() int
	GetMessageByLang(lang Lang) string
	Code() int
	Key() string
	TraceID() string
	SetTraceID()
}

// customErr struct
type customErr struct {
	// HTTP status code
	httpStatuscode int
	// Code of error
	code int
	// Key to translate
	key string
	// Message of error
	message string
	// trace id of error
	traceID string
}

// Error return error message
func (ce customErr) Error() string {
	return ce.message
}

// HTTPCode return http error code
func (ce customErr) HTTPCode() int {
	return ce.httpStatuscode
}

// GetMessageByLang return message with language
func (ce customErr) GetMessageByLang(lang Lang) string {
	mapOfLangErr, ok := mapOfLang[ce.key]
	if !ok {
		return ce.message
	}

	mess, ok := mapOfLangErr[lang]
	if !ok {
		return ce.message
	}

	return mess
}

// TraceID get id of trace error
func (ce customErr) TraceID() string {
	return ce.traceID
}

func (ce *customErr) SetTraceID() {
	val := make([]rune, 23)
	for i := range val {
		val[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	ce.traceID = string(val)
}

func (ce customErr) Code() int {
	return ce.code
}

func (ce customErr) Key() string {
	return ce.key
}

// NewCustomErrByMsg in constructor of CustomErr
func NewCustomErrByMsg(msg string, optArr ...Option) CustomErrorer {
	opts := convertOptArrayToOptions(optArr...)
	opt := opts.ToOption()
	return &customErr{
		message:        msg,
		key:            opt.Key,
		code:           opt.Code,
		httpStatuscode: opt.HTTPCode,
	}
}

// NewCustomErrByKey in constructor of CustomErr
func NewCustomErrByKey(key string, optArr ...Option) CustomErrorer {
	opts := convertOptArrayToOptions(optArr...)
	opt := opts.ToOption()
	return &customErr{
		code:           opt.Code,
		key:            key,
		message:        getMessageFromKey(key),
		httpStatuscode: opt.HTTPCode,
	}
}

func getMessageFromKey(key string) string {
	mapEnErr, ok := mapOfLang[key]
	if !ok {
		return ""
	}

	mess, ok := mapEnErr[EN]
	if !ok {
		return ""
	}

	return mess
}
