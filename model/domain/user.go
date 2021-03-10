package domain

import "github.com/phungvandat/source-template/utils/ctxkey"

// User holds information's user
type User struct {
	Base
	Name string `json:"name"`
}

// Context key of user ID
const (
	CtxKeyUserID ctxkey.CtxKey = "ctx_key_user_id"
)
