package routes

import (
	"github.com/danyducky/social/api/controllers"
	"github.com/danyducky/social/app"
)

type SignRoutes struct {
	context        app.Context
	signController controllers.SignController
}

func NewSignRoutes(
	context app.Context,
	signController controllers.SignController,
) SignRoutes {
	return SignRoutes{
		context:        context,
		signController: signController,
	}
}

func (s SignRoutes) Setup() {
	api := s.context.ApiGroup.Group("/sign")
	{
		api.POST("/", s.signController.Login)
		api.DELETE("/", s.signController.Logout)
	}
}
