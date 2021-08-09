package types

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	SessionToken string `json:"session_token"`
}

type GetUserRequest struct {
	Id int64 `json:"id"`
}
