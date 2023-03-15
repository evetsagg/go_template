package main

import (
	"go_template/src/database"
	"go_template/src/logger"

	"github.com/magiconair/properties"
)

type LoggingI interface {
	//todo add more functionality
	Info(string)
	Error(error)
	Debug(string)
	Fatal(error)
}

func main() {

	loggingService := logger.NewLogger()
	loggingService.Info("Initialized Logger")
	db := database.GetDatabase(loggingService)
	//db := database.GetMysqlDatabase(loggingService)
	loggingService.Info("Initialized Database")

	//Setup Dependency Injections and getting app properties
	productDao := database.NewProductDao(loggingService, db)
	prop := properties.MustLoadFile("config/app.properties", properties.UTF8)
	processor := NewProcessor(loggingService, prop, productDao)

	//Starting the Server to listen to request
	server := NewServer(loggingService, prop, processor)
	server.handleRequests()
	//todo set configuration for api clients if applicable

}
