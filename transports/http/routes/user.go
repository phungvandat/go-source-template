package routes

import (
	"github.com/phungvandat/source-template/endpoints"
	fwk "github.com/phungvandat/source-template/pkg/framework/http"
	jsonDec "github.com/phungvandat/source-template/transports/http/decode/json/user"
	jsonEnc "github.com/phungvandat/source-template/transports/http/encode/json"
)

// UserRoutes define all route relate to user
func UserRoutes(
	endpoints endpoints.Endpoints,
) func(r fwk.Router) {
	return func(r fwk.Router) {
		// swagger:route POST /users/login User Login
		// User login.
		// parameters:
		// + name: body
		//   in: body
		//	 type: UserLoginHTTPReq
		// responses:
		//	default: body:ErrorResponse
		// 	200: body:UserLoginHTTPRes
		r.POST("/login", fwk.NewHandle(
			jsonDec.LoginRequest,
			endpoints.User.Login,
			jsonEnc.EncodeJSONResponse,
		))
	}
}
