package middleware

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	userMod "server-golang/models/database"
	"time"
)

type jwtClaim struct {
	UserId uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJwt(data userMod.User) string {
	claims := &jwtClaim{
		data.Id,
		data.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	result, _ := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	return result
}
