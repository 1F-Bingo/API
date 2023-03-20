package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `gorm:"primaryKey"`
	Password   string
	BingoPicks []UserBingoPick
}
