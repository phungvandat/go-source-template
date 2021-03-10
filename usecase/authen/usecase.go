package authen

import (
	"github.com/go-redis/redis/v8"
	dbRedis "github.com/phungvandat/source-template/utils/config/db/redis"
	"github.com/phungvandat/source-template/utils/errs"
)

// UseCase interface
type UseCase interface {
	CreateToken(userID string) (accessToken, refreshToken string, err error)
}

type authen struct {
	rClient   *redis.Client
	jwtSecret []byte
	eTracer   errs.ErrTracer
}

// NewAuthenUseCase constructor of authen
func NewAuthenUseCase(jwtSecret string, eTracer errs.ErrTracer) UseCase {
	return &authen{
		rClient:   dbRedis.GetDB(),
		jwtSecret: []byte(jwtSecret),
		eTracer:   eTracer,
	}
}
