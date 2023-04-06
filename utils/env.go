package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables defined in the .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
