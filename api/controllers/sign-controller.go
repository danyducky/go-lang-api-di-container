package controllers

import (
	"net/http"

	"github.com/danyducky/social/domain/commands/sign"
	"github.com/gin-gonic/gin"
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
// @Success 200 {object} sign.LoginCommand
// @Param LoginCommand body sign.LoginCommand true "Login object"
// @Router /sign [post]
func (c *signController) Login(ctx *gin.Context) {
	var command sign.LoginCommand

	if err := ctx.BindJSON(&command); err != nil {
		return
	}

	ctx.JSON(http.StatusOK, command)
}

// Logout
// @Summary Allows the user to logout
// @Tags sign
// @Accept json
// @Produce json
// @Success 200 {object} app.Response
// @Router /sign [delete]
func (c *signController) Logout(ctx *gin.Context) {
	response := sign.LoginCommand{
		Email:    "morikov2000@gmail.com",
		Password: "123145212512",
		Age:      21,
	}
	ctx.JSON(http.StatusOK, response)
}
