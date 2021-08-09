package user

import (
	"github.com/allen012694/usersystem/models/profile_picture"
	"github.com/allen012694/usersystem/types"
	"gorm.io/gorm"
)

func GetUserByUsername(tx *gorm.DB, username string) (*User, error) {
	var user *User
	err := tx.Model(&User{}).Preload("ProfilePictures", func(q *gorm.DB) *gorm.DB {
		return q.Where(&profile_picture.ProfilePicture{Status: 1}).Order("updated_at DESC").Order("id DESC")
	}).Where(&User{Status: 1, Username: username}).First(&user).Error
	return user, err
}

func GetUserById(tx *gorm.DB, id int64) (*User, error) {
	var user *User
	err := tx.Model(&User{}).Preload("ProfilePictures", func(q *gorm.DB) *gorm.DB {
		return q.Where(&profile_picture.ProfilePicture{Status: 1}).Order("updated_at DESC").Order("id DESC")
	}).Where(&User{Status: 1}).First(&user, id).Error
	return user, err
}

func UpdateUserById(tx *gorm.DB, data *types.UpdateUserRequest) error {
	return tx.Where(&User{Id: data.UserId}).Updates(data).Error
}
