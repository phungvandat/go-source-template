package domain

import "github.com/phungvandat/source-template/utils/ctxkey"

// User holds information's user
type User struct {
	Base
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"-"`
}

// Context key of user ID
const (
	CtxKeyUserID         ctxkey.CtxKey = "ctx_key_user_id"
	UserAccessSessionID                = "user_access_session_id"
	UserRefreshSessionID               = "user_refresh_session_id"
)
