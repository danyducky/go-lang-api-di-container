package main

import (
	"context"

	"github.com/danyducky/social/api/controllers"
	"github.com/danyducky/social/api/handlers"
	"github.com/danyducky/social/api/middlewares"
	"github.com/danyducky/social/api/routes"
	"github.com/danyducky/social/app"
	docs "github.com/danyducky/social/docs"
	"github.com/danyducky/social/internal/common"
	"github.com/danyducky/social/internal/repositories"
	"github.com/danyducky/social/internal/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
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

// Main application container for dependency injection.
var Container = fx.Options(
	controllers.Container,
	repositories.Container,
	services.Container,
	handlers.Container,
	app.Container,
	middlewares.Container,
	routes.Container,
	common.Container,

	// Calls application lifecycle method.
	fx.Invoke(configure),
)

func main() {
	fx.New(Container).Run()
}

// Configures application lifecycle.
func configure(
	lifecycle fx.Lifecycle,
	handler app.Context,
	config app.Config,
	db app.Database,
	middlewares middlewares.Middlewares,
	routes routes.Routes,
) {
	conn, _ := db.Connection.DB()

	lifecycle.Append(fx.Hook{

		// Called when application is started.
		OnStart: func(context.Context) error {
			conn.SetMaxOpenConns(10)

			go func() {
				setEnviroment(config)

				// base path for all swagger [@Router] values.
				docs.SwaggerInfo.BasePath = "/api"

				// configure swagger support.
				handler.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

				middlewares.Setup()
				routes.Setup()

				handler.Gin.Run(":" + config.Application.Port)
			}()
			return nil
		},

		// Called when application is stopped.
		OnStop: func(context.Context) error {
			conn.Close()
			return nil
		},
	})
}

func setEnviroment(config app.Config) {
	if config.Environment.Mode == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
