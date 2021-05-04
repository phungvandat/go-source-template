package user

import (
	"context"

	"github.com/go-redis/redis/v8"
	iom "github.com/phungvandat/source-template/model/service/user"
	"github.com/phungvandat/source-template/pkg/errpkg"
	"github.com/phungvandat/source-template/usecase"
	"gorm.io/gorm"
)

// Service is user service
type Service interface {
	Login(ctx context.Context, in *iom.LoginIn) (*iom.LoginOut, error)
}

type svc struct {
	uc      usecase.Usecase
	eTracer errpkg.ErrTracer
	pg      *gorm.DB
	rd      *redis.Client
}

// NewService is constructor of user service
func NewService(pg *gorm.DB, rd *redis.Client, eTracer errpkg.ErrTracer, uc usecase.Usecase) Service {
	return &svc{
		pg:      pg,
		uc:      uc,
		eTracer: eTracer,
		rd:      rd,
	}
}
