package models

import (
	"gorm.io/gorm"
)

// Role entity.
type Role struct {
	gorm.Model

	// Role name.
	Name string

	// Users associated with current role.
	Users []*User `gorm:"many2many:user_roles;"`
}
