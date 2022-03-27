package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"site.com/config"
	"site.com/models"
)

type Connection struct {
	Host     string
	Port     uint16
	User     string
	Database string
	Password string
}

func (c *Connection) GetString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", c.Host, c.User, c.Password, c.Database, c.Port)
}

var (
	connection = Connection{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Database: "social",
		Password: "1",
	}
)

func Connect() *gorm.DB {
	config.CreateDatabaseIfNotExists(connection.Database)
	connection := postgres.Open(connection.GetString())
	database, err := gorm.Open(connection, &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// runs auto migration for given models.
	database.AutoMigrate(&models.User{}, &models.Role{})

	return database
}

func Close(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Cannot close database connection!")
	}

	dbSQL.Close()
}
