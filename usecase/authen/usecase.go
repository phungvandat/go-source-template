package authen

import (
	"github.com/go-redis/redis/v8"
	"github.com/phungvandat/source-template/model/domain"
	"github.com/phungvandat/source-template/pkg/errs"
	dbRedis "github.com/phungvandat/source-template/utils/config/db/redis"
)

// UseCase interface
type UseCase interface {
	CreateToken(userID domain.ID) (*createTokenRes, error)
	private()
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

func (authen) private() {
	// Anti-tampering
}
