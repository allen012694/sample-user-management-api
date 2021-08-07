package controllers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/allen012694/usersystem/config"
	"github.com/allen012694/usersystem/models/user"
	"github.com/allen012694/usersystem/types"
	"github.com/allen012694/usersystem/utils"
	"github.com/dgrijalva/jwt-go"
)

func Login(ctx context.Context, req *types.LoginRequest) (*types.LoginResponse, error) {
	// Look for existed user
	user, err := user.GetUserByUsername(req.Username)
	if err != nil {
		return nil, errors.New(config.ErrorUserNotExisted)
	}

	// Calculate password hash
	pwdHash := utils.SHA256WithSalt(req.Password, config.SALT)
	if user.PasswordHash != pwdHash {
		return nil, errors.New(config.ErrorPasswordWrong)
	}

	// Generate JWT token
	tokenizer := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        fmt.Sprint(user.Id),
		ExpiresAt: time.Now().Add(7 * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
	})
	token, err := tokenizer.SignedString([]byte(config.SECRET))
	if err != nil {
		return nil, err
	}

	// Store session into redis
	// TODO

	return &types.LoginResponse{SessionToken: token}, nil
}
