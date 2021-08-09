package user

import (
	"github.com/allen012694/usersystem/common"
	"github.com/allen012694/usersystem/types"
)

func GetUserByUsername(username string) (*User, error) {
	var user *User
	err := common.GetDB().Where(&User{Status: 1, Username: username}).First(&user).Error
	return user, err
}

func GetUserById(id int64) (*User, error) {
	var user *User
	err := common.GetDB().Where(&User{Status: 1}).First(&user, id).Error
	return user, err
}

func UpdateUserById(data *types.UpdateUserRequest) error {
	err := common.GetDB().Where(&User{Id: data.UserId}).Updates(data).Debug().Error
	return err
}
