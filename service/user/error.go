package user

import "github.com/phungvandat/source-template/pkg/errs"

// Special errors belong to user
var (
	ErrUsernameRequired = errs.NewCustomErrByCode(6)
)
