package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/allen012694/usersystem/config"
	"github.com/dgrijalva/jwt-go"
)

const EXPIRE_DURATION = 7 * 60 // minutes

type JwtTokenizer struct {
	secret   string
	signAlgo jwt.SigningMethod
}

type JwtClaims struct {
	Payload JwtPayload `json:"payload"`
	jwt.StandardClaims
}

type JwtPayload struct {
	Id int64 `json:"id"`
}

func NewJwtTokenizer() *JwtTokenizer {
	// default signing algorithm useing HS256
	return &JwtTokenizer{
		secret:   config.SECRET,
		signAlgo: jwt.SigningMethodHS256,
	}
}

func (tokenizer *JwtTokenizer) Generate(payload JwtPayload) (string, error) {
	token := jwt.NewWithClaims(tokenizer.signAlgo, JwtClaims{
		payload,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(EXPIRE_DURATION * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	return token.SignedString([]byte(tokenizer.secret))
}

func (tokenizer *JwtTokenizer) Validate(token string) (*JwtPayload, error) {
	res, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(tk *jwt.Token) (interface{}, error) {
		return []byte(tokenizer.secret), nil
	})
	if err != nil || !res.Valid {
		return nil, errors.New(config.ErrorLoginSessionInvalid)
	}
	claims, ok := res.Claims.(*JwtClaims)
	if !ok {
		return nil, errors.New(config.ErrorLoginSessionInvalid)
	}

	return &claims.Payload, nil
}

func ExtractJwtTokenFromHeaderString(authoirzationHeader string) (string, error) {
	parts := strings.Split(authoirzationHeader, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", errors.New(config.ErrorLoginSessionInvalid)
	}
	return parts[1], nil
}
