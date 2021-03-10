package jwtutil

import "github.com/phungvandat/source-template/utils/errs"

// jwtutil error
var (
	ErrUnexpectedTokenSigningMethod = errs.NewCustomErrByCode(2)
	ErrTokenInvalid                 = errs.NewCustomErrByCode(3)
)
