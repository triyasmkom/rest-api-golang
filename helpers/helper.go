package helpers

import (
	"github.com/joho/godotenv"
	"os"
)

func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return ":8000"
	}

	return ":" + port
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed load env file")
	}
}
