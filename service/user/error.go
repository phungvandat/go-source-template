package user

import "github.com/phungvandat/source-template/utils/errs"

// Special errors belong to user
var (
	ErrUsernameRequired = errs.NewCustomErrByCode(6)
)
