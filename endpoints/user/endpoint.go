package user

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/phungvandat/source-template/service"
)

// Endpoint of user
type Endpoint struct {
	Login endpoint.Endpoint
}

// NewEndpoint create user endpoint
func NewEndpoint(s service.Service) Endpoint {
	return Endpoint{
		Login: makeLoginEndpoint(s.UserSvc),
	}
}
