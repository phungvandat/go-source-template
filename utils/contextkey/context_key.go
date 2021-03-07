package contextkey

import "context"

// CtxKey use for context
type CtxKey int

const (
	// CtxKeyUserID define key send through context
	CtxKeyUserID CtxKey = iota
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
