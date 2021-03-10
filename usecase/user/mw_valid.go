package user

import "github.com/phungvandat/source-template/utils/errs"

type validMw struct {
	UseCase
	eTracer errs.ErrTracer
}

// NewValidationMiddleware is constructor of validation middleware
func NewValidationMiddleware(eTracer errs.ErrTracer) func(UseCase) UseCase {
	return func(next UseCase) UseCase {
		return &validMw{
			UseCase: next,
			eTracer: eTracer,
		}
	}
}
