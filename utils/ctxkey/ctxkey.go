package ctxkey

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// ctxKey use for context
type ctxKey string

// List of context key send through context
const (
	ctxKeyUserID  ctxKey = "ctx_key_user_id"
	ctxKeyLang    ctxKey = "ctx_key_lang"
	ctxKeyPGDB    ctxKey = "ctx_key_pg_db"
	ctxKeyRedisDB ctxKey = "ctx_key_redis_db"
)

// GetStrValue get string value with key
func getStrValue(ctx context.Context, key ctxKey) string {
	var (
		val       = ctx.Value(key)
		valStr, _ = val.(string)
	)
	return valStr
}

// SetUserID assign user ID to context
func SetUserID(parent context.Context, userID string) context.Context {
	return context.WithValue(parent, ctxKeyUserID, userID)
}

// GetUserID get user ID from context
func GetUserID(ctx context.Context) string {
	return getStrValue(ctx, ctxKeyUserID)
}

// SetLang assign language to context
func SetLang(parent context.Context, userID string) context.Context {
	return context.WithValue(parent, ctxKeyLang, userID)
}

// GetLang get language from context
func GetLang(ctx context.Context) string {
	return getStrValue(ctx, ctxKeyLang)
}

// SetPG assign pg db to context
func SetPG(parent context.Context, db *gorm.DB) context.Context {
	return context.WithValue(parent, ctxKeyPGDB, db)
}

// GetPG get pg db from context
func GetPG(ctx context.Context) (*gorm.DB, error) {
	var (
		val       = ctx.Value(ctxKeyPGDB)
		valDB, ok = val.(*gorm.DB)
	)

	if !ok || valDB == nil {
		return nil, errors.New("missing assign pg to context")
	}

	return valDB, nil
}

// SetRedis assign redis db to context
func SetRedis(parent context.Context, client *redis.Client) context.Context {
	return context.WithValue(parent, ctxKeyRedisDB, client)
}

// GetRedis get redis db from context
func GetRedis(ctx context.Context) (*redis.Client, error) {
	var (
		val       = ctx.Value(ctxKeyRedisDB)
		valDB, ok = val.(*redis.Client)
	)

	if !ok || valDB == nil {
		return nil, errors.New("missing assign redis to context")
	}

	return valDB, nil
}
