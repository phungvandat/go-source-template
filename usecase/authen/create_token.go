package authen

import (
	"time"

	"github.com/phungvandat/source-template/model/domain"
	"github.com/phungvandat/source-template/utils/jwtutil"
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

func (uc *authen) CreateToken(userID domain.ID) (*createTokenRes, error) {
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

	accessToken, err := jwtutil.CreateToken(atGenData)
	if err != nil {
		return nil, uc.eTracer.Trace(err)
	}

	refreshToken, err := jwtutil.CreateToken(rfGenData)
	if err != nil {
		return nil, uc.eTracer.Trace(err)
	}

	return &createTokenRes{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		AccessSessionID:  atSessionID,
		RefreshSessionID: rfSessionID,
	}, nil
}
