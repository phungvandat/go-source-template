package user

import (
	"context"

	iom "github.com/phungvandat/source-template/model/service/user"
	"github.com/phungvandat/source-template/pkg/endpoint"
	"github.com/phungvandat/source-template/service/user"
	"github.com/phungvandat/source-template/utils/errs"
)

func makeLoginEndpoint(s user.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		in, ok := request.(*iom.LoginSvcIn)
		if !ok {
			return nil, errs.ErrSomethingWentWrong
		}
		return s.Login(ctx, in)
	}
}
