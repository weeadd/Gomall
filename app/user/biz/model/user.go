package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex;type:varchar(255) not null"`
	PasswordHashed string `gorm:"not null;type:varchar(255)"`
}

func (User) TableName() string {
	return "user"
}

func Create(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	return &user, err
}
