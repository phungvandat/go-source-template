package token

import (
	"context"

	"github.com/phungvandat/source-template/model/domain"
	"github.com/phungvandat/source-template/pkg/errpkg"
)

// UseCase interface
type UseCase interface {
	CreateToken(ctx context.Context, userID domain.ID) (accessToken, refreshToken string, err error)
	private()
}

type token struct {
	jwtSecret []byte
	eTracer   errpkg.ErrTracer
}

// NewTokenUseCase constructor of token
func NewTokenUseCase(jwtSecret string, eTracer errpkg.ErrTracer) UseCase {
	return &token{
		jwtSecret: []byte(jwtSecret),
		eTracer:   eTracer,
	}
}

func (token) private() {
	// Anti-tampering
}
