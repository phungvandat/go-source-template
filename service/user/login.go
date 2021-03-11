package user

import (
	"context"

	"github.com/phungvandat/source-template/model/domain"
	iom "github.com/phungvandat/source-template/model/service/user"
	"github.com/phungvandat/source-template/utils/errs"
	"gorm.io/gorm"
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
			return nil, errs.ErrUserNotFound
		}
		return nil, s.eTracer.Trace(err)
	}

	// TODO: compare with password
	ctRes, err := s.uc.Authen.CreateToken(user.ID)

	if err != nil {
		return nil, s.eTracer.Trace(err)
	}

	return &iom.LoginSvcOut{
		AccessToken:  ctRes.AccessSessionID,
		RefreshToken: ctRes.RefreshSessionID,
	}, nil
}
