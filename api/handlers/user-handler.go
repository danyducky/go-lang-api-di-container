package handlers

import (
	"github.com/danyducky/social/domain/commands/user"
	"github.com/danyducky/social/domain/models"
	"github.com/danyducky/social/internal/services"
	"github.com/dranikpg/dto-mapper"
)

type UserHandler interface {
	HandleRegister(command user.RegisterCommand) uint
}

type userHandler struct {
	userService services.UserService
	mapper      dto.Mapper
}

func NewUserHandler(
	userService services.UserService,
	mapper dto.Mapper,
) UserHandler {
	return &userHandler{
		userService: userService,
		mapper:      mapper,
	}
}

func (h userHandler) HandleRegister(command user.RegisterCommand) uint {
	var user models.User

	h.mapper.Map(&user, command)

	h.userService.Register(&user)

	return user.ID
}
