package errs

import (
	"net/http"
)

// Common errors
var (
	ErrMethodNotAllowed   = NewCustomErrByCode(http.StatusMethodNotAllowed)
	ErrSomethingWentWrong = NewCustomErrByCode(http.StatusInternalServerError)
	ErrNotFound           = NewCustomErrByCode(http.StatusNotFound)
	ErrBodyNotAllowed     = NewCustomErrByCode(1)
	ErrPermissionDenied   = NewCustomErrByCode(http.StatusForbidden)
)
