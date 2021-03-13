package jwtutil

import "github.com/phungvandat/source-template/pkg/errs"

// jwtutil error
var (
	ErrUnexpectedTokenSigningMethod = errs.NewCustomErrByCode(2)
	ErrTokenInvalid                 = errs.NewCustomErrByCode(3)
)
