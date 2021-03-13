package user

import (
	"context"

	iom "github.com/phungvandat/source-template/model/service/user"
	"github.com/phungvandat/source-template/utils/errs"
)

func (mw mwValid) Login(ctx context.Context, in *iom.LoginSvcIn) (*iom.LoginSvcOut, error) {
	if in.Username == "" {
		return nil, mw.eTracer.Trace(errs.ErrUsernameRequired)
	}

	return mw.Service.Login(ctx, in)
}
