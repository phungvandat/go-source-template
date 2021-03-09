package contextkey

import "context"

// CtxKey use for context
type CtxKey string

// List of context key send through context
const (
	CtxKeyUserID CtxKey = "ctx_key_user_id"
	CtxKeyLang   CtxKey = "ctx_key_lang"
)

// GetUserID get user ID from context
func GetUserID(ctx context.Context) string {
	var (
		ctxUserID = ctx.Value(CtxKeyUserID)
		userID, _ = ctxUserID.(string)
	)
	return userID
}

// SetUserID set user ID to context
func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, CtxKeyUserID, userID)
}

// SetLang set language to context
func SetLang(ctx context.Context, lang string) context.Context {
	return context.WithValue(ctx, CtxKeyLang, lang)
}

// GetLang get language from context
func GetLang(ctx context.Context) string {
	var (
		ctxLang = ctx.Value(CtxKeyLang)
		lang, _ = ctxLang.(string)
	)
	return lang
}
