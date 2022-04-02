package repositories

import (
	"fmt"

	"github.com/danyducky/social/app"
	"github.com/danyducky/social/domain/models"
)

// User repository implementation pattern.
type UserRepository interface {
	Get() []models.User
	Find(id int) models.User
	FindBy(field string, value interface{}) models.User
}

// User repository structure.
type userRepository struct {
	database app.Database
}

// Creates user repository instance.
func NewUserRepository(db app.Database) UserRepository {
	return &userRepository{
		database: db,
	}
}

// Returns a list of users.
func (r *userRepository) Get() []models.User {
	var users []models.User
	r.database.Connection.Find(&users)
	return users
}

// Returns the found user entity.
func (r *userRepository) Find(id int) models.User {
	var user models.User
	r.database.Connection.First(&user, id)
	return user
}

// Returns the found user by condition.
func (r *userRepository) FindBy(field string, value interface{}) models.User {
	var user models.User
	stmt := fmt.Sprintf("%s = ?", field)
	r.database.Connection.Where(stmt, value).First(&user)
	return user
}
