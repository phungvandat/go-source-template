package env

import (
	"os"
	"sync"
)

var (
	once         sync.Once
	isProduction bool
)

// IsProduction check production environment
func IsProduction() bool {
	once.Do(func() {
		isProduction = ENV() == "production"
	})

	return isProduction
}

// ENV func
func ENV() string {
	return os.Getenv("ENV")
}

var (
	httpPortOnce sync.Once
	httpPort     string
)

func HTTPPort() string {
	httpPortOnce.Do(func() {
		httpPort = os.Getenv("HTTP_PORT")
	})
	return httpPort
}

var (
	pgSourceOnce sync.Once
	pgSource     string
)

func PGSource() string {
	pgSourceOnce.Do(func() {
		pgSource = os.Getenv("PG_SOURCE")
	})
	return pgSource
}

var (
	redisSourceOnce sync.Once
	redisSource     string
)

func RedisSource() string {
	redisSourceOnce.Do(func() {
		redisSource = os.Getenv("REDID_SOURCE")
	})

	return redisSource
}

var (
	jwtSecretOne sync.Once
	jwtSecret    string
)

func JWTSecret() string {
	jwtSecretOne.Do(func() {
		jwtSecret = os.Getenv("JWT_SECRET")
	})
	return jwtSecret
}
