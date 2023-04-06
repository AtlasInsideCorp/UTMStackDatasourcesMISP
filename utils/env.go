package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables defined in the .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

// Getenv returns the environment variable
func Getenv(key string) string {
	value, defined := os.LookupEnv(key)
	if !defined {
		log.Fatalf("error loading environment variable: %s: environment variable does not exist\n", key)
	}
	if (value == "") || (value == " ") {
		log.Fatalf("error loading environment variable: %s: empty environment variable\n", key)
	}
	return value
}
