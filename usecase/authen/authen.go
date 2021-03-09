package authen

import (
	"github.com/go-redis/redis/v8"
	dbRedis "github.com/phungvandat/source-template/utils/config/db/redis"
)

type authen struct {
	rClient   *redis.Client
	jwtSecret []byte
}

// NewAuthenUseCase constructor of authen
func NewAuthenUseCase(jwtSecret string) UseCase {
	return &authen{
		rClient:   dbRedis.GetDB(),
		jwtSecret: []byte(jwtSecret),
	}
}

func (uc *authen) CreateToken() {

}
