package domain

// User holds information's user
type User struct {
	Base
	CompanyID ID `json:"company_id"`

	Fullname string   `json:"fullname"`
	Username string   `json:"username"`
	Password string   `json:"-"`
	RoleType RoleType `json:"role_type"`
}

type RoleType int

const (
	RoleTypeIncompetent RoleType = 0
	RoleTypeLimited     RoleType = 1
	RoleTypeUnlimited   RoleType = 2
)

func (rt RoleType) Int() int {
	return int(rt)
}

const (
	UserAccessSessionID  = "user_access_session_id"
	UserRefreshSessionID = "user_refresh_session_id"
)
