package database

import (
	"go_template/src/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type LoggingI interface {
	//todo add more functionality
	Info(string)
	Error(error)
	Debug(string)
	Fatal(error)
}

func GetDatabase(logger LoggingI) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		logger.Fatal(err)
	}
	db.AutoMigrate(&model.Product{})

	return db
}
