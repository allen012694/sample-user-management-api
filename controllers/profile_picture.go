package controllers

import (
	"context"

	"github.com/allen012694/usersystem/common"
	"github.com/allen012694/usersystem/models/profile_picture"
	"github.com/allen012694/usersystem/types"
)

func AddProfilePicture(ctx context.Context, req *types.AddProfilePictureRequest) (*profile_picture.ProfilePicture, error) {
	tx := common.GetDB().Begin()

	// deactivate current picture
	if err := profile_picture.DeactivateAllPictures(tx, req.UserId); err != nil {
		tx.Rollback()
		return nil, err
	}

	// create new picture with status activated
	profilePic, err := profile_picture.CreateProfilePicture(tx, req)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return profilePic, nil
}
