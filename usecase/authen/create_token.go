package authen

import (
	"time"

	"github.com/phungvandat/source-template/utils/jwtutil"
	uuid "github.com/satori/go.uuid"
)

const (
	userIDKey    = "user_id"
	sessionIDKey = "session_id"
)

func (uc *authen) CreateToken(userID string) (string, string, error) {
	var (
		// Access token
		atExp       = time.Hour * 12
		atSessionID = uuid.NewV4().String()
		atGenData   = jwtutil.TokenInfo{
			MapClaimsData: map[string]interface{}{
				userIDKey:    userID,
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
				userIDKey:    userID,
				sessionIDKey: rfSessionID,
			},
			Secret:      uc.jwtSecret,
			ExpiredTime: rfExp,
		}
	)

	accessToken, err := jwtutil.CreateToken(atGenData)
	if err != nil {
		return "", "", uc.eTracer.Trace(err)
	}

	refreshToken, err := jwtutil.CreateToken(rfGenData)
	if err != nil {
		return "", "", uc.eTracer.Trace(err)
	}

	return accessToken, refreshToken, nil
}
