package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Key is not found in .env file")
	}

	return os.Getenv(key)
}

func SetConfig(key string, value string) error {
	return os.Setenv(key, value)
}
