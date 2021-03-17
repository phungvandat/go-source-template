package token

import (
	"context"
	"fmt"
	"time"

	"github.com/phungvandat/source-template/model/domain"
	"github.com/phungvandat/source-template/utils/ctxkey"
	"github.com/phungvandat/source-template/utils/errs"
	"github.com/phungvandat/source-template/utils/jwtutil"
	"github.com/phungvandat/source-template/utils/redisutil"
	uuid "github.com/satori/go.uuid"
)

const (
	userIDKey    = "user_id"
	sessionIDKey = "session_id"
)

type createTokenRes struct {
	AccessToken      string
	RefreshToken     string
	AccessSessionID  string
	RefreshSessionID string
}

func (uc *token) CreateToken(ctx context.Context, userID domain.ID) (accessToken, refreshToken string, err error) {
	var (
		userIDStr = userID.String()
		// Access token
		atExp       = time.Hour * 12
		atSessionID = uuid.NewV4().String()
		atGenData   = jwtutil.TokenInfo{
			MapClaimsData: map[string]interface{}{
				userIDKey:    userIDStr,
				sessionIDKey: atSessionID,
			},
			Secret:      uc.jwtSecret,
			ExpiredTime: atExp,
		}
		// Refresh token
		rfExp       = time.Hour * 24 * 7
		rfSessionID = uuid.NewV4().String()
		rfGenData   = jwtutil.TokenInfo{
			MapClaimsData: map[string]interface{}{
				userIDKey:    userIDStr,
				sessionIDKey: rfSessionID,
			},
			Secret:      uc.jwtSecret,
			ExpiredTime: rfExp,
		}
	)

	accessToken, err = jwtutil.CreateToken(atGenData)
	if err != nil {
		return "", "", uc.eTracer.Trace(errs.ErrSomethingWentWrong, err)
	}

	refreshToken, err = jwtutil.CreateToken(rfGenData)
	if err != nil {
		return "", "", uc.eTracer.Trace(errs.ErrSomethingWentWrong, err)
	}

	err = uc.setSessionIDToRedis(ctx, userIDStr, atSessionID, rfSessionID)
	if err != nil {
		return "", "", uc.eTracer.Trace(errs.ErrSomethingWentWrong, err)
	}

	return accessToken, refreshToken, nil
}

func (uc *token) setSessionIDToRedis(ctx context.Context, userID, accessTokenID, refreshTokenID string) error {
	rClient, err := ctxkey.GetRedis(ctx)
	if err != nil {
		return uc.eTracer.Trace(errs.ErrSomethingWentWrong, err)
	}

	redisAccessKey := redisTokenByUserID(domain.UserAccessSessionID, userID, accessTokenID)
	err = redisutil.SetWithKey(ctx, rClient, redisAccessKey, userID, 12*time.Hour)
	if err != nil {
		return uc.eTracer.Trace(errs.ErrSomethingWentWrong, err)
	}

	redisRefreshKey := redisTokenByUserID(domain.UserRefreshSessionID, userID, accessTokenID)
	err = redisutil.SetWithKey(ctx, rClient, redisRefreshKey, userID, 7*24*time.Hour)
	if err != nil {
		return uc.eTracer.Trace(errs.ErrSomethingWentWrong, err)
	}

	return nil
}

func redisTokenByUserID(prefix, userID, id string) string {
	return fmt.Sprintf("%v_%v_%v", prefix, userID, id)
}
