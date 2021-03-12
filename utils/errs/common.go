package errs

import (
	"net/http"
)

// Common errors
var (
	ErrMethodNotAllowed        = NewCustomErrByCode(http.StatusMethodNotAllowed, Option{HTTPStatusCode: http.StatusMethodNotAllowed})
	ErrSomethingWentWrong      = NewCustomErrByCode(http.StatusInternalServerError, Option{HTTPStatusCode: http.StatusInternalServerError})
	ErrNotFound                = NewCustomErrByCode(http.StatusNotFound, Option{HTTPStatusCode: http.StatusNotFound})
	ErrBodyNotAllowed          = NewCustomErrByCode(1, Option{HTTPStatusCode: http.StatusNotAcceptable})
	ErrPermissionDenied        = NewCustomErrByCode(http.StatusForbidden, Option{HTTPStatusCode: http.StatusForbidden})
	ErrIDIsInvalid             = NewCustomErrByCode(4, Option{HTTPStatusCode: http.StatusBadRequest})
	ErrInvalidTypeToDecodeJSON = NewCustomErrByCode(7, Option{HTTPStatusCode: http.StatusBadRequest})
	ErrRouteNotFound           = NewCustomErrByCode(8, Option{HTTPStatusCode: http.StatusNotFound})
)
