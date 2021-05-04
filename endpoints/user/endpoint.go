package user

import (
	"github.com/phungvandat/source-template/pkg/endpoint"
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
