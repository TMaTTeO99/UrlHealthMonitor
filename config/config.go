package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() (*ConfigData, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &ConfigData{
		API_KEY:             os.Getenv("VIRUSTOTAL_API_KEY"),
		DB_USER_NAME:        os.Getenv("DB_USER_NAME"),
		USER_PASSWD:         os.Getenv("USER_PASSWD"),
		VIRUSTOTAL_BASE_URL: os.Getenv("VIRUSTOTAL_BASE_URL"),
		PREFIX:              os.Getenv("PREFIX"),
		DB_HOST:             os.Getenv("DB_HOST"),
		DB_PORT:             os.Getenv("DB_PORT"),
		DB_NAME:             os.Getenv("DB_NAME"),
	}, nil

}
