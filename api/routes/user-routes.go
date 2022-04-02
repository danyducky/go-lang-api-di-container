package routes

import (
	"github.com/danyducky/social/api/controllers"
	"github.com/danyducky/social/app"
)

type UserRoutes struct {
	context        app.Context
	userController controllers.UserController
}

func NewUserRoutes(
	context app.Context,
	userController controllers.UserController,
) UserRoutes {
	return UserRoutes{
		context:        context,
		userController: userController,
	}
}

func (s UserRoutes) Setup() {
	api := s.context.ApiGroup.Group("/user")
	{
		api.GET("/", s.userController.GetMe)
		api.POST("/", s.userController.Register)
	}
}
