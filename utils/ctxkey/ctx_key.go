package ctxkey

import "context"

// CtxKey use for context
type CtxKey string

// GetStrValue get string value with key
func GetStrValue(ctx context.Context, key CtxKey) string {
	var (
		val       = ctx.Value(key)
		valStr, _ = val.(string)
	)
	return valStr
}
