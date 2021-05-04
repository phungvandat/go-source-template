package errs

import (
	"net/http"

	"github.com/phungvandat/source-template/pkg/errpkg"
)

// jwt token error
var (
	ErrTokenUnexpectedSigningMethod = errpkg.NewCustomErrByKey(
		"token.unexpected_signing_method",
		errpkg.OptHTTPCode(http.StatusBadRequest),
	)
	ErrTokenInvalid = errpkg.NewCustomErrByKey(
		"token.invalid",
		errpkg.OptHTTPCode(http.StatusBadRequest),
	)
)
