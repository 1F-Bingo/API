package models

import (
	"API/internal"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `gorm:"primaryKey"`
	Password   string
	BingoPicks []UserBingoPick
}

func GetOrError(username string) (User, error) {
	var user User
	internal.GetDB().First(&user, "username = ?", username)
	if user.Username == "" {
		return User{}, errors.New("invalid username")
	}
	return user, nil
}

func LoginCheck(user User, password string) (string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := internal.GenerateToken(user.Username)

	if err != nil {
		return "", err
	}

	return token, nil

}
