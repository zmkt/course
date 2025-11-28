package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadKey() (string, error) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	key := os.Getenv("KEY")
	return key, nil
}
