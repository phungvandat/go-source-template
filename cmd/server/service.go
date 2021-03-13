package main

import (
	"github.com/phungvandat/source-template/pkg/errs"
	"github.com/phungvandat/source-template/service"
	"github.com/phungvandat/source-template/service/user"
	"github.com/phungvandat/source-template/usecase"
)

func initService(
	eTracer errs.ErrTracer,
) service.Service {
	var (
		userSvc, _ = service.Compose(
			user.NewService(nil, nil, eTracer, usecase.Usecase{}),
			user.NewValidationMiddleware(eTracer),
		).(user.Service)

		svc = service.Service{
			UserSvc: userSvc,
		}
	)

	return svc
}
