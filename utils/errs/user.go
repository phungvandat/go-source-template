package errs

import (
	"net/http"

	"github.com/phungvandat/source-template/pkg/errpkg"
)

// Special errors belong to user
var (
	ErrUsernameRequired = errpkg.NewCustomErrByKey(
		"login.username_required",
		errpkg.OptHTTPCode(http.StatusBadRequest),
	)
	ErrUserNotFound = errpkg.NewCustomErrByKey(
		"user.not_found",
		errpkg.OptHTTPCode(http.StatusNotFound),
	)
	ErrWrongPass = errpkg.NewCustomErrByKey(
		"login.wrong_pass",
		errpkg.OptHTTPCode(http.StatusBadRequest),
	)
	ErrPassRequired = errpkg.NewCustomErrByKey(
		"login.pass_required",
		errpkg.OptHTTPCode(http.StatusBadRequest),
	)
)
