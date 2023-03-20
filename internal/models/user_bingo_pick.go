package models

import (
	"gorm.io/gorm"
)

type UserBingoPick struct {
	gorm.Model
	BoardPosition int16
	league        string
	UserUsername  string `gorm:"primaryKey"`
	BingoItemID   uint   `gorm:"primaryKey"`
}
