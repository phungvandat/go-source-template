package errpkg

// Option custom error option
type Option struct {
	Message  string
	Code     int
	IsTraced bool
	HTTPCode int
	Key      string
}

// Options is list of Option
type Options []Option

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

		if val.HTTPCode > 0 {
			opt.HTTPCode = val.HTTPCode
		}

		if val.Message != "" {
			opt.Message = val.Message
		}

		opt.IsTraced = val.IsTraced

	}
	return opt
}

func OptMessage(mess string) Option {
	return Option{Message: mess}
}

func OptCode(code int) Option {
	return Option{Code: code}
}

func OptIsTraced(isTraced bool) Option {
	return Option{IsTraced: true}
}

func OptHTTPCode(code int) Option {
	return Option{HTTPCode: code}
}

func OptKey(key string) Option {
	return Option{Key: key}
}
