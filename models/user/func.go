package user

import "github.com/allen012694/usersystem/context"

func GetUserByUsername(username string) (*User, error) {
	var user *User
	err := context.GetDB().Where(&User{Username: username}).First(&user).Error
	return user, err
}
