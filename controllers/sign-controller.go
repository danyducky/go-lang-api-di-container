package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	sign "site.com/commands"
	helper "site.com/infrastructure"
)

type SignController interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type signController struct {
}

func NewSignController() SignController {
	return &signController{}
}

// Login
// @Summary Allows the user to login
// @Tags sign
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Param LoginCommand body sign.LoginCommand true "Login object"
// @Router /sign [post]
func (c *signController) Login(ctx *gin.Context) {
	var command sign.LoginCommand
	var response helper.Response

	if err := ctx.BindJSON(&command); err != nil {
		response = helper.BadResponse("Something went wrong.", err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response = helper.OkResponse("Everything is ok. Here the user must login :)", command)
	ctx.JSON(http.StatusOK, response)
}

// Logout
// @Summary Allows the user to logout
// @Tags sign
// @Accept json
// @Produce json
// @Success 200 {object} helper.Response
// @Router /sign [delete]
func (c *signController) Logout(ctx *gin.Context) {
	response := helper.EmptyResponse("Everything is ok. Here the user must logout :)")
	ctx.JSON(http.StatusOK, response)
}
