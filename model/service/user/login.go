package user

// LoginSvcIn define input data for login handle
type LoginSvcIn struct {
	Username string `json:"user_name"`
}

// LoginSvcOut define output data for login handle
type LoginSvcOut struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
