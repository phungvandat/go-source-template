package user

import (
	"context"

	"github.com/phungvandat/source-template/model/domain"
	iom "github.com/phungvandat/source-template/model/service/user"
	"gorm.io/gorm"

	"github.com/phungvandat/source-template/utils/errs"
)

func (s *svc) Login(ctx context.Context, in *iom.LoginSvcIn) (*iom.LoginSvcOut, error) {
	var (
		username = in.Username
		user     = &domain.User{}
		err      error
	)

	err = s.pg.Where("username = ?", username).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, s.eTracer.Trace(errs.ErrUserNotFound)
		}
		return nil, s.eTracer.Trace(err)
	}

	// TODO: compare with password
	accessToken, refreshToken, err := s.uc.Token.CreateToken(ctx, user.ID)
	if err != nil {
		return nil, s.eTracer.Trace(err)
	}

	return &iom.LoginSvcOut{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
