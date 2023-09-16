package helpers

import (
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return ":8081"
	}

	return ":" + port
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed load env file")
	}
}

func HashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func VerifyPassword(password string, hashPassword string) bool {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), hash)
	if err != nil {
		return false
	}
	return true
}
