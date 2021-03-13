package errs

// ErrLang type
type ErrLang string

type langError struct {
	Message string `json:"message"`
}

// List of error language
const (
	VN ErrLang = "vn"
	EN ErrLang = "en" // Default language
)

// CustomErrorer interface
type CustomErrorer interface {
	Error() string
	HTTPStatusCode() int
	GetMessageByLang(lang ErrLang) string
	IsTraced() bool
	SetTraced(val bool)
	Code() int
}

// customErr struct
type customErr struct {
	// HTTP status code
	httpStatuscode int
	// Code to translate
	code int
	// Message of error
	message string
	// Is error traced?
	isTraced bool
}

// Error return error message
func (ce customErr) Error() string {
	return ce.message
}

// HTTPStatusCode return http error code
func (ce customErr) HTTPStatusCode() int {
	return ce.httpStatuscode
}

// GetMessageByLang return message with language
func (ce customErr) GetMessageByLang(lang ErrLang) string {
	mapOfLangErr, ok := mapOfLang[lang]
	if !ok {
		return ce.message
	}

	mess, ok := mapOfLangErr[ce.code]
	if !ok {
		return ce.message
	}

	return mess.Message
}

// IsTraced check error traced
func (ce customErr) IsTraced() bool {
	return ce.isTraced
}

func (ce *customErr) SetTraced(val bool) {
	ce.isTraced = val
}

func (ce customErr) Code() int {
	return ce.code
}

// Option custom error option
type Option struct {
	Message        string
	Code           int
	IsTraced       bool
	HTTPStatusCode int
}

// Options is list of Option
type Options []Option

// NewCustomErrByMsg in constructor of CustomErr
func NewCustomErrByMsg(msg string, optArr ...Option) CustomErrorer {
	opts := convertOptArrayToOptions(optArr...)
	opt := opts.ToOption()
	return &customErr{
		message:        msg,
		code:           opt.Code,
		httpStatuscode: opt.HTTPStatusCode,
		isTraced:       opt.IsTraced,
	}
}

// NewCustomErrByCode in constructor of CustomErr
func NewCustomErrByCode(code int, optArr ...Option) CustomErrorer {
	opts := convertOptArrayToOptions(optArr...)
	opt := opts.ToOption()
	return &customErr{
		code:           code,
		message:        getMessageFromCode(code),
		httpStatuscode: opt.HTTPStatusCode,
		isTraced:       opt.IsTraced,
	}
}

func convertOptArrayToOptions(arr ...Option) Options {
	var opts = make(Options, len(arr))
	for idx := range arr {
		opts[idx] = arr[idx]
	}
	return opts
}

// ToOption convert list of option to option
func (opts Options) ToOption() Option {
	var opt Option
	for idx := range opts {
		var val = opts[idx]
		if val.Code > 0 {
			opt.Code = val.Code
		}

		if val.HTTPStatusCode > 0 {
			opt.HTTPStatusCode = val.HTTPStatusCode
		}

		if val.Message != "" {
			opt.Message = val.Message
		}

		opt.IsTraced = val.IsTraced

	}
	return opt
}

func getMessageFromCode(code int) string {
	mapEnErr, ok := mapOfLang[EN]
	if !ok {
		return ""
	}
	mess, ok := mapEnErr[code]
	if !ok {
		return ""
	}
	return mess.Message
}
