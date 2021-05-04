package jwtutil

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/phungvandat/source-template/utils/errs"
)

// TokenInfo holds data to create jwt token
type TokenInfo struct {
	MapClaimsData map[string]interface{}
	ExpiredTime   time.Duration
	Secret        []byte
}

// CreateToken create jwt token
func CreateToken(ti TokenInfo) (string, error) {
	var (
		claims        = ti.convertToJWTClaims()
		token         = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, err = token.SignedString(ti.Secret)
	)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (ti TokenInfo) convertToJWTClaims() jwt.MapClaims {
	var mapClaims = jwt.MapClaims(ti.MapClaimsData)

	mapClaims["exp"] = time.Now().Add(ti.ExpiredTime).Unix()
	return mapClaims
}

// VerifyToken verify token
func VerifyToken(tokenStr string, secret []byte) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.ErrTokenUnexpectedSigningMethod
		}

		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errs.ErrTokenInvalid
	}

	return claims, nil
}
