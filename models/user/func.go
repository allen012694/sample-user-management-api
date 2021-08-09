package user

import "github.com/allen012694/usersystem/common"

func GetUserByUsername(username string) (*User, error) {
	var user *User
	err := common.GetDB().Where(&User{Username: username}).First(&user).Error
	return user, err
}

func GetUserById(id int64) (*User, error) {
	var user *User
	err := common.GetDB().First(&user, id).Error
	return user, err
}
