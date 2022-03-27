package config

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

var (
	// Representing data for connect to default postgres database.
	defaultConnection = pgx.ConnConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Database: "postgres",
		Password: "1",
	}
)

// Allows to create database if not exists.
func CreateDatabaseIfNotExists(dbName string) {
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
