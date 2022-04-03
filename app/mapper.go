package app

import (
	"github.com/danyducky/social/domain/commands/user"
	"github.com/danyducky/social/domain/models"
	"github.com/danyducky/social/internal/common"
	"github.com/dranikpg/dto-mapper"
)

func NewMapper(hashService common.HashService) dto.Mapper {
	mapper := dto.Mapper{}

	mapper.AddInspectFunc(func(user *models.User, command user.RegisterCommand) {
		user.PasswordHash, _ = hashService.Create(command.Password)
	})

	return mapper
}
