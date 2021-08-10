package types

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserId       int64  `json:"-"`
	SessionToken string `json:"session_token"`
}

type GetUserRequest struct {
	UserId int64 `json:"user_id"`
}

type UpdateUserRequest struct {
	UserId int64 `json:"user_id" gorm:"-"` // identify prop, not data patch

	Nickname *string `json:"nickname" gorm:"column:nickname;"`
}

func (UpdateUserRequest) TableName() string {
	return "user_tab"
}
