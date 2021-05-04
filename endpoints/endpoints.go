package endpoints

import (
	"github.com/phungvandat/source-template/endpoints/user"
	"github.com/phungvandat/source-template/service"
)

// Endpoints holds all endpoint
type Endpoints struct {
	User user.Endpoint
}

func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		User: user.NewEndpoint(s),
	}
}
