package main

import (
	"ml-x-men/config"
	"ml-x-men/internal/adapters/db"
	server "ml-x-men/internal/adapters/http/gin.server"
	"ml-x-men/internal/adapters/logger"
	"ml-x-men/internal/application"
)

func main() {
	appConfig := config.New()
	ginServer := server.NewServer(8082, server.DebugMode)

	routerLogger := logger.NewLogger("TEST",
		"DEBUG",
		"text",
	)
	storage, err := db.New(appConfig.DBConfig)
	if err != nil {
		routerLogger.Log(err, "Ups")
		panic(err)
	}
	server.NewRouterWithLogger(application.HandlerConstructor{
		Logger:  routerLogger,
		Storage: storage,
	}.New(),
		routerLogger).SetRoutes(ginServer.Router)

	ginServer.Start()

}
