package user

import (
	"context"

	"github.com/phungvandat/source-template/model/domain"
	iom "github.com/phungvandat/source-template/model/service/user"
)

func (uc *user) FindOne(ctx context.Context, in *iom.FindOneSvcIn) (*iom.FindOneSvcOut, error) {
	var (
		user = &domain.User{}
	)

	err := uc.pg.Find(user).Error

	if err != nil {
		return nil, uc.eTracer.Trace(err)
	}

	return &iom.FindOneSvcOut{User: user}, nil
}
