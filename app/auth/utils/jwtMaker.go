package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var stSignKey = []byte("getttoken")

type JwtCustomClaims struct {
	UserId int32
	jwt.RegisteredClaims
}

func GenerateToken(UserId int32) (string, error) {
	iJwtCustomClaims := JwtCustomClaims{
		UserId: UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustomClaims)
	return token.SignedString(stSignKey)
}

func ParseToken(tokenStr string) (JwtCustomClaims, error) {
	iJwtCustomClaims := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return stSignKey, nil
	})

	if err != nil && !token.Valid {
		err = errors.New("invalid token")
	}
	return iJwtCustomClaims, err
}
