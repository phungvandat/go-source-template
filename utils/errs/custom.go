package errs

// CustomErr struct
type CustomErr struct {
	code     int
	message  string
	isTraced bool
}

// Error return error message
func (err CustomErr) Error() string {
	return err.message
}

// StatusCode return error code
func (err CustomErr) StatusCode() int {
	return err.code
}

// NewCustomErr in constructor of CustomErr
func NewCustomErr(msg string, code int) CustomErr {
	return CustomErr{
		message: msg,
		code:    code,
	}
}
