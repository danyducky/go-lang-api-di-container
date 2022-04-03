package services

import (
	"github.com/danyducky/social/domain/models"
	"github.com/danyducky/social/internal/repositories"
)

type UserService interface {
	Register(user *models.User)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s userService) Register(user *models.User) {
	s.userRepository.Insert(user)
}
