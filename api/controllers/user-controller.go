package controllers

import (
	"net/http"

	"github.com/danyducky/social/api/dtos"
	"github.com/danyducky/social/api/handlers"
	"github.com/danyducky/social/domain/commands/user"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetMe(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type userController struct {
	handler handlers.UserHandler
}

func NewUserController(handler handlers.UserHandler) UserController {
	return &userController{
		handler: handler,
	}
}

// GetMe
// @Summary Represent information about current user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} dtos.UserDto
// @Router /user [get]
func (c *userController) GetMe(ctx *gin.Context) {
	data := dtos.UserDto{
		Firstname: "Danil",
		Lastname:  "Morikov",
	}

	ctx.JSON(http.StatusOK, data)
}

// Register
// @Summary Allows user to register
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} user.RegisterCommand
// @Param RegisterCommand body user.RegisterCommand true "Register object"
// @Router /user [post]
func (c *userController) Register(ctx *gin.Context) {
	var command user.RegisterCommand

	if err := ctx.BindJSON(&command); err != nil {
		return
	}
	userId := c.handler.HandleRegister(command)

	ctx.JSON(http.StatusOK, userId)
}
