package database

import (
	"go_template/src/model"

	"gorm.io/driver/mysql"

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

// Sqlite
func GetDatabase(logger LoggingI) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		logger.Fatal(err)
	}
	db.AutoMigrate(&model.Product{})

	return db
}

// Mysql
func GetMysqlDatabase(logger LoggingI) *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Fatal(err)
	}
	db.AutoMigrate(&model.Product{})

	return db
}
