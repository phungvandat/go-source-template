package user

import "github.com/phungvandat/source-template/pkg/errpkg"

type mwValid struct {
	Service
	eTracer errpkg.ErrTracer
}

// NewValidationMiddleware is constructor of validation middleware
func NewValidationMiddleware(eTracer errpkg.ErrTracer) func(Service) Service {
	return func(next Service) Service {
		return &mwValid{
			eTracer: eTracer,
			Service: next,
		}
	}
}
