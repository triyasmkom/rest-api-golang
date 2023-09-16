package helpers

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"os"
	"server-golang/models/response"
	"strconv"
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

func VerifyPassword(password string, hashPassword string) (response.Response, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return response.Response{}, err
	}
	return response.Response{}, nil
}

func Debug() bool {
	var debug, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		fmt.Println("Error Debug Variabel: ", err)
		return false
	}
	return debug
}
