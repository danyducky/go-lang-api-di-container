package models

import (
	"time"

	"gorm.io/gorm"
)

// User entity.
type User struct {
	gorm.Model

	// User email.
	Email string

	// Hashed user password.
	PasswordHash string

	// User firstname.
	Firstname string

	// User lastname.
	Lastname string

	// User phone number.
	PhoneNumber string

	// User gender.
	Gender UserGender

	// Last login date.
	LastLogin time.Time

	// User roles.
	Roles []*Role `gorm:"many2many:user_roles;"`
}

// Represents user gender.
type UserGender uint8

const (
	Male   UserGender = 1
	Female UserGender = 2
)
