package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"server-golang/models/database"
	"server-golang/models/response"
	"time"
)

func GenerateJwt(data database.User) (response.Token, error) {
	iat := time.Now()
	exp := time.Now().Add(time.Hour * 72)
	claims := jwt.MapClaims{
		"email":    data.Email,
		"username": data.Username,
		"exp":      exp.Unix(),
		"iat":      iat.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := "secret"
	result, err := token.SignedString([]byte(secretKey))

	if err != nil {
		fmt.Println("Error Generate Jwt: ", err)
		return response.Token{}, err
	}
	return response.Token{
		Token:   result,
		Created: iat.String(),
		Expired: exp.String(),
	}, nil
}
