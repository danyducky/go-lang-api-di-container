package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"site.com/dtos"
	helper "site.com/infrastructure"
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
// @Router /user/me [get]
func (c *userController) GetMe(ctx *gin.Context) {
	data := dtos.UserDto{
		Firstname: "Danil",
		Lastname:  "Morikov",
	}
	response := helper.OkResponse("This endpoint will give information about the user.", data)
	ctx.JSON(http.StatusOK, response)
}
