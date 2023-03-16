package app

import (
	"r-G7D/go_gin_restful/domains"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DefDB() {
	db, err := gorm.Open(sqlite.Open("go_gin_restful.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//TODO: db polling
	db.AutoMigrate(&domains.Driver{})

	DB = db
}
