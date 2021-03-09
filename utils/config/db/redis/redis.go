package redis

import (
	"context"
	"strconv"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/phungvandat/source-template/utils/stringutil"
)

var (
	once    = sync.Once{}
	redisDB *redis.Client
)

// InitRedisConn create a redis connection
func InitRedisConn(connStr string) {
	once.Do(func() {
		var (
			addr    = stringutil.GetStringInBetween(connStr, "addr=", " ")
			pass    = stringutil.GetStringInBetween(connStr, "pass=", " ")
			dbStr   = stringutil.GetStringInBetween(connStr, "db=", " ")
			db, err = strconv.Atoi(dbStr)
		)
		if err != nil {
			panic(err)
		}

		redisDB = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pass,
			DB:       db,
		})
		if _, err := redisDB.Ping(context.Background()).Result(); err != nil {
			panic(err)
		}
	})
}

// GetDB return a db connection
func GetDB() *redis.Client {
	if redisDB == nil {
		panic("missed InitRedisConn")
	}
	return redisDB
}
