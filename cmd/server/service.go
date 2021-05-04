package main

import (
	"github.com/phungvandat/source-template/config/db/pg"
	"github.com/phungvandat/source-template/config/db/redis"
	"github.com/phungvandat/source-template/config/env"
	"github.com/phungvandat/source-template/pkg/errpkg"
	"github.com/phungvandat/source-template/service"
	"github.com/phungvandat/source-template/service/user"
	"github.com/phungvandat/source-template/usecase"
	"github.com/phungvandat/source-template/usecase/token"
)

func initService(
	eTracer errpkg.ErrTracer,
) service.Service {
	var (
		usecase = usecase.Usecase{
			Token: token.NewTokenUseCase(env.JWTSecret(), eTracer),
		}
		userSvc, _ = service.Compose(
			user.NewService(pg.GetDB(), redis.GetDB(), eTracer, usecase),
			user.NewValidationMiddleware(eTracer),
		).(user.Service)

		svc = service.Service{
			UserSvc: userSvc,
		}
	)

	return svc
}
