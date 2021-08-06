package models

import "time"

type ProfilePicture struct {
	Id         int       `json:"id" gorm:"column:id;"`
	UserId     int       `json:"user_id" gorm:"column:user_id;"`
	PictureUrl string    `json:"picture_url" gorm:"column:picture_url;"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at;"`
	Status     int       `json:"status" gorm:"column:status;"`
}

func (ProfilePicture) TableName() string {
	return "profile_picture_tab"
}
