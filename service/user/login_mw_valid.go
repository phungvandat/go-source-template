package user

import (
	"context"

	iom "github.com/phungvandat/source-template/model/service/user"
	"github.com/phungvandat/source-template/utils/errs"
)

func (mw mwValid) Login(ctx context.Context, in *iom.LoginIn) (*iom.LoginOut, error) {
	if in.Username == "" {
		return nil, mw.eTracer.Trace(errs.ErrUsernameRequired)
	}

	if in.Password == "" {
		return nil, mw.eTracer.Trace(errs.ErrPassRequired)
	}

	return mw.Service.Login(ctx, in)
}
