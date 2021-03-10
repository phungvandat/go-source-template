package ctxkey

import "context"

// CtxKey use for context
type CtxKey string

// List of context key send through context
const (
	CtxKeyUserID CtxKey = "ctx_key_user_id"
	CtxKeyLang   CtxKey = "ctx_key_lang"
)

// GetStrValue get string value with key
func GetStrValue(ctx context.Context, key CtxKey) string {
	var (
		val       = ctx.Value(CtxKeyUserID)
		valStr, _ = val.(string)
	)
	return valStr
}
