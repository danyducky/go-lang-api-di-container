package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"site.com/config"
	"site.com/database"
	docs "site.com/docs"
)

// https://github.com/swaggo/swag/blob/master/example/celler/main.go
// Source code where represent swagger example.

// @title           Social API.
// @version         1.0
// @description     This is a social web server.

// @contact.name   We are united
// @contact.url    https://oursharedlink.com
// @contact.email  ourworkemail@gmail.com

// @BasePath  /api

var (
	db *gorm.DB = database.Connect()
)

func main() {
	// close database connection when the application shutdown.
	defer database.Close(db)

	// gin engine initializtion.
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api"

	// setup request/response middlewares.
	config.SetupMiddlewares(router)

	// setup endpoint routes.
	config.SetupRouter(router)

	// configure swagger support.
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Listen and Server in 0.0.0.0:5001
	router.Run(":5001")
}
