package middleware

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"server-golang/models/database"
	"time"
)

type jwtClaim struct {
	UserId uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJwt(data database.User) (string, error) {
	claims := &jwtClaim{
		data.Id,
		data.Email,
		jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	result, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil{
		return "", err
	}
	return result, nil
}
