package user

import (
	"context"
	"time"

	"github.com/phungvandat/source-template/model/domain"
	iom "github.com/phungvandat/source-template/model/service/user"
	"gorm.io/gorm"

	"github.com/phungvandat/source-template/utils/errs"
	"github.com/phungvandat/source-template/utils/redisutil"
)

func (s *svc) Login(ctx context.Context, in *iom.LoginSvcIn) (*iom.LoginSvcOut, error) {
	return nil, nil
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

	err = s.setSessionIDToRedis(ctx, ctRes.AccessSessionID, ctRes.RefreshSessionID)
	if err != nil {
		return nil, s.eTracer.Trace(err)
	}

	return &iom.LoginSvcOut{
		AccessToken:  ctRes.AccessToken,
		RefreshToken: ctRes.RefreshToken,
	}, nil
}

func (s *svc) setSessionIDToRedis(ctx context.Context, accessTokenID, refreshTokenID string) error {
	err := redisutil.SetWithKey(ctx, s.rd, domain.UserAccessSessionID, accessTokenID, 12*time.Hour)
	if err != nil {
		return s.eTracer.Trace(err)
	}

	err = redisutil.SetWithKey(ctx, s.rd, domain.UserRefreshSessionID, refreshTokenID, 7*24*time.Hour)
	if err != nil {
		return s.eTracer.Trace(err)
	}

	return nil
}
