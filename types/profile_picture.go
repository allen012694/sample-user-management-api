package types

type AddProfilePictureRequest struct {
	UserId     int64  `json:"user_id" `
	PictureUrl string `json:"pickture_url"`
}
