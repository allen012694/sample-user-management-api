package controllers

import (
	"context"
	"errors"

	"github.com/allen012694/usersystem/common"
	"github.com/allen012694/usersystem/config"
	"github.com/allen012694/usersystem/models/user"
	"github.com/allen012694/usersystem/types"
	"github.com/allen012694/usersystem/utils"
)

func Login(ctx context.Context, req *types.LoginRequest) (*types.LoginResponse, error) {
	// Look for existed user
	user, err := user.GetUserByUsername(common.GetDB(), req.Username)
	if err != nil {
		return nil, errors.New(config.ErrorUserNotExisted)
	}

	// Calculate password hash
	pwdHash := utils.SHA256WithSalt(req.Password, config.SALT)
	if user.PasswordHash != pwdHash {
		return nil, errors.New(config.ErrorPasswordWrong)
	}

	// Generate JWT token
	tokenizer := utils.NewJwtTokenizer()
	token, err := tokenizer.Generate(utils.JwtPayload{Id: int64(user.Id)})
	if err != nil {
		return nil, err
	}

	// Store session into redis
	utils.PutStoreSession(ctx, token)

	return &types.LoginResponse{SessionToken: token}, nil
}

func GetUser(ctx context.Context, req *types.GetUserRequest) (*user.User, error) {
	user, err := user.GetUserById(common.GetDB(), req.UserId)
	if err != nil {
		return nil, errors.New(config.ErrorUserNotExisted)
	}

	return user, err
}

func UpdateUser(ctx context.Context, req *types.UpdateUserRequest) (*user.User, error) {
	err := user.UpdateUserById(common.GetDB(), req)
	if err != nil {
		return nil, err
	}

	user, err := user.GetUserById(common.GetDB(), req.UserId)
	return user, err
}
