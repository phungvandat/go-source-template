package user

import (
	"context"

	iom "github.com/phungvandat/source-template/model/service/user"
	"github.com/phungvandat/source-template/utils/errs"
	"gorm.io/gorm"
)

// UseCase of user
type UseCase interface {
	FindOne(ctx context.Context, in *iom.FindOneSvcIn) (*iom.FindOneSvcOut, error)
}

type user struct {
	pg      *gorm.DB
	eTracer errs.ErrTracer
}

// NewUseCase is constructor of user usecase
func NewUseCase(pg *gorm.DB, eTracer errs.ErrTracer) UseCase {
	return &user{
		pg:      pg,
		eTracer: eTracer,
	}
}
