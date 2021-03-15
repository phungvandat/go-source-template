package redis

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/phungvandat/source-template/config/env"
)

// CreateTestDatabase create to redis connection to testing
func CreateTestDatabase(t *testing.T) (*redis.Client, func()) {
	connStr := "addr=localhost:6378 pass=password db=0"
	if env.RedisTestingSource() != "" {
		connStr = env.RedisTestingSource()
	}

	addr, pass, db := parseRedisConfig(connStr)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})
	ctx := context.Background()
	if _, err := client.Ping(ctx).Result(); err != nil {
		t.Fatal(err)
	}

	return client, func() {
		if err := client.FlushAll(ctx).Err(); err != nil {
			t.Fatal(err)
		}
		if err := client.Close(); err != nil {
			t.Fatal(err)
		}
	}
}
