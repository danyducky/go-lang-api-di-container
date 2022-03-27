package models

import (
	"gorm.io/gorm"
)

// User entity.
type User struct {
	gorm.Model

	// User firstname.
	Firstname string

	// User lastname.
	Lastname string

	// User roles.
	Roles []*Role `gorm:"many2many:user_roles;"`
}
