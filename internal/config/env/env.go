package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic("Error load .env file")
	}

	if key == "" {
		panic("A key should be provider.")
	}

	return os.Getenv(key)
}
