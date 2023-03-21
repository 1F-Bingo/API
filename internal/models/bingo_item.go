package models

import (
	"gorm.io/gorm"
	"time"
)

type BingoItem struct {
	gorm.Model
	Description     string `gorm:"index"`
	RequireCount    int64
	OccurrenceCount int64
	FinishedAt      time.Time
	UsersPicked     []UserBingoPick
	League          string // If League is global, anyone can use it
}

func (bi BingoItem) HasOccurred() bool {
	return bi.RequireCount == bi.OccurrenceCount
}
