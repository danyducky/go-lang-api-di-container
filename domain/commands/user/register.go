package user

import "github.com/danyducky/social/domain/models"

// Represent the fields that allows the user to register.
type RegisterCommand struct {
	Email       string            `json:"email" binding:"required,email"`
	Password    string            `json:"password" binding:"required,min=6"`
	Firstname   string            `json:"firstname" binding:"required"`
	Lastname    string            `json:"lastname" binding:"required"`
	Gender      models.UserGender `json:"gender" binding:"required"`
	PhoneNumber string            `json:"phoneNumber" binding:"required,e164"`
}
