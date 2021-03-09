package jwtutil

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type TokenInfo struct {
	MapClaimsData map[string]interface{}
	ExpiredTime   time.Duration
	Secret        []byte
}

func CreateToken(info TokenInfo) {

}

func (ti TokenInfo) convertToJWTClaims() jwt.MapClaims {
	var mapClaims = make(jwt.MapClaims)

	for key := range ti.MapClaimsData {
		mapClaims[key] = ti.MapClaimsData[key]
	}

	mapClaims["exp"] = time.Now().Add(ti.ExpiredTime).Unix()
	return mapClaims
}
