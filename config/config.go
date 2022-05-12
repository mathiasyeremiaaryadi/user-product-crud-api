package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig(key string) string {
	err := godotenv.Load("env/.env.dev")
	if err != nil {
		fmt.Print("Cannot load the environment file")
	}
	return os.Getenv(key)
}