package user

import (
	"context"

	"github.com/phungvandat/source-template/model/domain"
	iom "github.com/phungvandat/source-template/model/service/user"
	"gorm.io/gorm"

	cryptopgk "github.com/phungvandat/source-template/pkg/crypto"
	"github.com/phungvandat/source-template/utils/ctxkey"
	"github.com/phungvandat/source-template/utils/errs"
)

func (s *svc) Login(ctx context.Context, in *iom.LoginIn) (*iom.LoginOut, error) {
	var (
		username = in.Username
		pass     = in.Password
		user     = &domain.User{}
		err      error
	)

	err = s.pg.Where("username = ?", username).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, s.eTracer.Trace(errs.ErrUserNotFound)
		}
		return nil, s.eTracer.Trace(errs.ErrSomethingWentWrong, err)
	}

	passMatched, err := cryptopgk.CompareValue(pass, user.Password)
	if !passMatched || err != nil {
		return nil, s.eTracer.Trace(errs.ErrWrongPass, err)
	}

	accessToken, refreshToken, err := s.uc.Token.CreateToken(ctxkey.SetRedis(ctx, s.rd), user.ID)
	if err != nil {
		return nil, s.eTracer.Trace(err)
	}

	return &iom.LoginOut{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
