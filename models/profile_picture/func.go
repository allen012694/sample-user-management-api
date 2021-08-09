package profile_picture

import (
	"github.com/allen012694/usersystem/types"
	"gorm.io/gorm"
)

func DeactivateAllPictures(tx *gorm.DB, userId int64) error {
	return tx.Model(&ProfilePicture{}).Where(&ProfilePicture{UserId: userId, Status: 1}).Update("status", 0).Error
}

func CreateProfilePicture(tx *gorm.DB, createReq *types.AddProfilePictureRequest) (*ProfilePicture, error) {
	profilePic := &ProfilePicture{
		UserId:     createReq.UserId,
		PictureUrl: createReq.PictureUrl,
		Status:     1,
	}
	err := tx.Create(&profilePic).Error
	return profilePic, err
}
