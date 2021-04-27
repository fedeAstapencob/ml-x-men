package main

import (
	server "ml-x-men/internal/adapters/http/gin.server"
	"ml-x-men/internal/adapters/logger"
	"ml-x-men/internal/application"
)

func main() {
	ginServer := server.NewServer(8082, server.DebugMode)

	routerLogger := logger.NewLogger("TEST",
		"DEBUG",
		"text",
	)
	server.NewRouterWithLogger(application.HandlerConstructor{
		Logger: routerLogger,
	}.New(),
		routerLogger).SetRoutes(ginServer.Router)

	ginServer.Start()

}
