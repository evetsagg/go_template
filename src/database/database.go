package database

import (
	"go_template/src/logger"
	"go_template/src/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDatabase(logger logger.LoggingI) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		logger.Fatal(err)
	}
	db.AutoMigrate(&model.Product{})
	
	return db
}