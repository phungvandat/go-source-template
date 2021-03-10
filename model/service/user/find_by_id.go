package user

import "github.com/phungvandat/source-template/model/domain"

// FindOneSvcIn define input data for handle
type FindOneSvcIn struct {
	// ID of user
	ID domain.ID `json:"id"`
	// Username of user
	Username string `json:"username"`
}

// FindOneSvcOut define output data for handle
type FindOneSvcOut struct {
	User *domain.User `json:"user"`
}
