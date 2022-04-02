package app

import (
	"fmt"
	"os"

	"github.com/danyducky/social/domain/models"
	"github.com/jackc/pgx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Application database wrapper.
type Database struct {
	Connection *gorm.DB
}

// Creates application database instance.
func NewDatabase(config Config) Database {
	createDatabaseIfNotExists("social")
	connection := postgres.Open(config.ConnectionString)
	database, err := gorm.Open(connection, &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// runs auto migration for given models.
	database.AutoMigrate(&models.User{}, &models.Role{})

	return Database{
		Connection: database,
	}
}

var (
	// Represents data for connect to default postgres database.
	defaultConnection = pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Database: "postgres",
		Password: "1",
	}
)

// Allows to create database if not exists.
func createDatabaseIfNotExists(dbName string) {
	conn, err := pgx.Connect(defaultConnection)

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Checking if database exists. Create, if not.
	rows, err := conn.Query("SELECT * FROM pg_database WHERE datname = $1", dbName)
	if !rows.Next() {
		createDatabase(conn, dbName)
	}
	defer rows.Close()
}

// Create database command.
func createDatabase(conn *pgx.Conn, dbName string) {
	stmt := fmt.Sprintf("CREATE DATABASE %s", dbName)
	tag, err := conn.Exec(stmt)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stdout, tag.RowsAffected())
}
