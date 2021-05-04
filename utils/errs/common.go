package errs

import (
	"net/http"

	"github.com/phungvandat/source-template/pkg/errpkg"
)

// Common errors
var (
	ErrMethodNotAllowed = errpkg.NewCustomErrByMsg(
		"method not allowed",
		errpkg.OptHTTPCode(http.StatusMethodNotAllowed),
	)
	ErrSomethingWentWrong = errpkg.NewCustomErrByMsg(
		"something went wrong",
		errpkg.OptHTTPCode(http.StatusInternalServerError),
	)
	ErrNotFound = errpkg.NewCustomErrByMsg(
		"not found",
		errpkg.OptHTTPCode(http.StatusNotFound),
	)
	ErrBodyNotAllowed = errpkg.NewCustomErrByKey(
		"data.body_not_allowed",
		errpkg.OptHTTPCode(http.StatusNotAcceptable),
	)
	ErrPermissionDenied = errpkg.NewCustomErrByMsg(
		"permission denied",
		errpkg.OptHTTPCode(http.StatusForbidden),
	)
	ErrRouteNotFound = errpkg.NewCustomErrByMsg(
		"route not found",
		errpkg.OptHTTPCode(http.StatusNotFound),
	)
)
