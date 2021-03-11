package user

import (
	"context"

	iom "github.com/phungvandat/source-template/model/service/user"
)

func (mw mwValid) Login(ctx context.Context, in *iom.LoginSvcIn) (*iom.LoginSvcOut, error) {
	if in.Username == "" {
		return nil, mw.eTracer.Trace(ErrUsernameRequired)
	}

	return mw.Service.Login(ctx, in)
}
