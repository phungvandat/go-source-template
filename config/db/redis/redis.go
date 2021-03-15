package redis

import (
	"context"
	"strconv"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/phungvandat/source-template/utils/helper"
	"github.com/phungvandat/source-template/utils/logger"
)

var (
	once    = sync.Once{}
	redisDB *redis.Client
)

// InitRedisConn create a redis connection
func InitRedisConn(connStr string) {
	once.Do(func() {
		addr, pass, db := parseRedisConfig(connStr)

		redisDB = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pass,
			DB:       db,
		})
		if _, err := redisDB.Ping(context.Background()).Result(); err != nil {
			panic(err)
		}
		logger.Info("redis db connected")
	})
}

func parseRedisConfig(connStr string) (addr, pass string, db int) {
	addr = helper.GetStringInBetween(connStr, "addr=", " ")
	pass = helper.GetStringInBetween(connStr, "pass=", " ")
	dbStr := helper.GetStringInBetween(connStr, "db=", " ")
	db, err := strconv.Atoi(dbStr)

	if err != nil {
		panic(err)
	}

	return addr, pass, db
}

// GetDB return a db connection
func GetDB() *redis.Client {
	if redisDB == nil {
		panic("missed InitRedisConn")
	}
	return redisDB
}

func Close() {
	if redisDB == nil {
		panic("missed InitRedisConn")
	}

	err := redisDB.Close()
	if err != nil {
		logger.Error("failed close redis connection by error: %v", err)
	}

	logger.Info("redis db closed")
}
