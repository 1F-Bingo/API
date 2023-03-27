package internal

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	var dbErr error
	if db == nil {
		db, dbErr = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	}
	if dbErr != nil {
		log.Fatal("db connection error:", dbErr)
		panic(dbErr)
	}

	return db
}
