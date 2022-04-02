package app

import (
	"encoding/json"
	"os"
)

// Application configuration.
type Config struct {
	Environment struct {
		Mode string `json:"mode"`
	} `json:"environment"`
	Application struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"application"`
	ConnectionString string `json:"connectionString"`
}

// Creates configuration instance based on ./appsettings.json file.
func NewConfig() Config {
	var config Config
	file, err := os.Open("appsettings.json")
	defer file.Close()
	if err != nil {
		panic("Appsettings.json not found, please fix it.")
	}
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&config)
	return config
}
