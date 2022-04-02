package controllers

import (
	"net/http"

	"github.com/danyducky/social/api/dtos"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetMe(ctx *gin.Context)
}

type userController struct {
}

func NewUserController() UserController {
	return &userController{}
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
