package user

import (
	"context"

	iom "github.com/phungvandat/source-template/model/service/user"
	"github.com/phungvandat/source-template/pkg/endpoint"
	"github.com/phungvandat/source-template/service/user"
)

func makeLoginEndpoint(s user.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.Login(ctx, request.(*iom.LoginIn))
	}
}
