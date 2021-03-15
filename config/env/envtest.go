package env

import (
	"os"
	"sync"
)

var (
	tPGSourceOnce sync.Once
	tPGSource     string
)

func PGTestingSource() string {
	tPGSourceOnce.Do(func() {
		tPGSource = os.Getenv("PG_TESTING_SOURCE")
	})
	return tPGSource
}

var (
	tRedisSourceOnce sync.Once
	tRedisSource     string
)

func RedisTestingSource() string {
	tRedisSourceOnce.Do(func() {
		tRedisSource = os.Getenv("REDID_TESTING_SOURCE")
	})

	return tRedisSource
}
