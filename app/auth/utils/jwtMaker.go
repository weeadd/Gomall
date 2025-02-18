package utils

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

func Macke(UserId int32) (token string, err error) {
	customClaim, err := NewCustomClaims(UserId)
	if err != nil {
		return "", err
	}
	then := jwt.NewWithClaims(jwt.SigningMethodES256, customClaim)
	token, err = then.SignedString([]byte("gettoken"))
	return
}

func secret() jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		return []byte("gettoken"), nil
	}
}

func ParseToken(token string) (bool, error) {
	tokn, err := jwt.Parse(token, secret())
	if err != nil {
		return false, errors.New("Fail to parse token")
	}
	if !tokn.Valid {
		return false, errors.New("Token is invalid")
	}
	return true, nil
}
