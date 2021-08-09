package user

import (
	"time"

	"github.com/allen012694/usersystem/models/profile_picture"
)

type User struct {
	Id           int64     `json:"id" gorm:"column:id;"`
	Username     string    `json:"username" gorm:"column:username;"`
	PasswordHash string    `json:"-" gorm:"column:password_hash;"`
	Nickname     string    `json:"nickname" gorm:"column:nickname;"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at;"`
	Status       int       `json:"status" gorm:"column:status;"`

	// reference
	ProfilePictures []*profile_picture.ProfilePicture `json:"profile_pictures"`
}

func (User) TableName() string {
	return "user_tab"
}
