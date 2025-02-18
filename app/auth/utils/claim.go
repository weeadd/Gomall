package utils

import (
	"errors"
	"time"
)

type CustomClaims struct {
	UserID    int32
	IssuedAt  time.Time
	ExpiredAt time.Time
}

func NewCustomClaims(UserID int32) (*CustomClaims, error) {
	customClaim := &CustomClaims{
		UserID:    UserID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Minute * 5),
	}
	return customClaim, nil
}

func (c *CustomClaims) Valid() error {
	if time.Now().After(c.ExpiredAt) {
		return errors.New("Token has expired")
	}
	return nil
}
