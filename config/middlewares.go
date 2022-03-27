package config

import (
	"github.com/gin-gonic/gin"
	"site.com/middlewares"
)

func SetupMiddlewares(router *gin.Engine) {
	// sequence of application middlewares.
	// Order matters!
	sequence := []func() gin.HandlerFunc{
		middlewares.JsonMiddleware,
	}

	// register all application middlewares in for loop.
	for _, middleware := range sequence {
		router.Use(middleware())
	}
}
