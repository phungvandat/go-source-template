package token

import (
	"context"
	"time"

	"github.com/phungvandat/source-template/model/domain"
	"github.com/phungvandat/source-template/utils/ctxkey"
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
		return "", "", uc.eTracer.Trace(err)
	}

	refreshToken, err = jwtutil.CreateToken(rfGenData)
	if err != nil {
		return "", "", uc.eTracer.Trace(err)
	}

	err = uc.setSessionIDToRedis(ctx, atSessionID, rfSessionID)
	if err != nil {
		return "", "", uc.eTracer.Trace(err)
	}

	return accessToken, refreshToken, nil
}

func (uc *token) setSessionIDToRedis(ctx context.Context, accessTokenID, refreshTokenID string) error {
	rClient, err := ctxkey.GetRedis(ctx)
	if err != nil {
		return uc.eTracer.Trace(err)
	}

	err = redisutil.SetWithKey(ctx, rClient, domain.UserAccessSessionID, accessTokenID, 12*time.Hour)
	if err != nil {
		return uc.eTracer.Trace(err)
	}

	err = redisutil.SetWithKey(ctx, rClient, domain.UserRefreshSessionID, refreshTokenID, 7*24*time.Hour)
	if err != nil {
		return uc.eTracer.Trace(err)
	}

	return nil
}
