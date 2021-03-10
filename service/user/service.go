package user

import (
	"github.com/phungvandat/source-template/usecase"
	"github.com/phungvandat/source-template/utils/errs"
)

// Service is user service
type Service struct {
	uc      usecase.Usecase
	eTracer errs.ErrTracer
}

// NewService is constructor of user service
func NewService(uc usecase.Usecase, eTracer errs.ErrTracer) Service {
	return Service{
		uc:      uc,
		eTracer: eTracer,
	}
}
