package models

import (
	"gorm.io/gorm"
)

type BingoItem struct {
	gorm.Model
	Description     string `gorm:"index"`
	RequireCount    int64
	OccurrenceCount int64
	UsersPicked     []UserBingoPick
}

func (bi BingoItem) HasOccurred() bool {
	return bi.RequireCount == bi.OccurrenceCount
}
