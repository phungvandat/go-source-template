package user

// swagger:model UserLoginHTTPReq
// LoginIn define input data for login handle
type LoginIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// swagger:model UserLoginHTTPRes
// LoginOut define output data for login handle
type LoginOut struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
