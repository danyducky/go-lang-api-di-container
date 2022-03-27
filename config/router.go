package config

import (
	"github.com/gin-gonic/gin"
	"site.com/controllers"
)

var (
	signController = controllers.NewSignController()
	userController = controllers.NewUserController()
)

// Setup route groups for gin engine.
// All controllers routes are placed here.
func SetupRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		signGroup(api)
		userGroup(api)
	}
}

// Sign controller routes group.
func signGroup(api *gin.RouterGroup) {
	sign := api.Group("/sign")
	{
		sign.POST("/", signController.Login)
		sign.DELETE("/", signController.Logout)
	}
}

// User controller routes group.
func userGroup(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		user.GET("/me", userController.GetMe)
	}
}
