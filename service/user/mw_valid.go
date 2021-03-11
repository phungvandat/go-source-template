package user

import "github.com/phungvandat/source-template/utils/errs"

type mwValid struct {
	Service
	eTracer errs.ErrTracer
}

// NewValidationMiddleware is constructor of validation middleware
func NewValidationMiddleware(eTracer errs.ErrTracer) func(Service) Service {
	return func(next Service) Service {
		return &mwValid{
			eTracer: eTracer,
		}
	}
}
