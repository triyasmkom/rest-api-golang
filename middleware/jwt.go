package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	helper "server-golang/helpers"
	"server-golang/models/database"
	"server-golang/models/response"
	"strconv"
	"time"
)

func GenerateJwt(data database.User) (response.Token, error) {
	secretKey := os.Getenv("JWT_KEY")
	jwtExp := os.Getenv("JWT_EXP")
	expTime, errExp := strconv.Atoi(jwtExp)
	if errExp != nil {
		if helper.Debug() {
			fmt.Println("Error Generate Jwt: ", errExp)
		}
		expTime = 5 // exp 5 hour if env null
	}

	iat := time.Now()
	exp := time.Now().Add(time.Hour * time.Duration(expTime))
	claims := jwt.MapClaims{
		"id":       data.Id,
		"role":     data.Roles,
		"email":    data.Email,
		"username": data.Username,
		"exp":      exp.Unix(),
		"iat":      iat.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(secretKey))

	if err != nil {
		if helper.Debug() {
			fmt.Println("Error Generate Jwt: ", err)
		}
		return response.Token{}, err
	}
	return response.Token{
		Token:   result,
		Created: iat.String(),
		Expired: exp.String(),
	}, nil
}

func VerifyJwt(tokenString string) response.Response {
	secretKey := []byte(os.Getenv("JWT_KEY"))

	// Untuk melakukan validasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validasi algoritma tanda tangan
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signature not valid")
		}
		return secretKey, nil
	})

	// Handle error saat parsing token
	if err != nil {
		if helper.Debug() {
			fmt.Println("Error Verify Jwt:", err)
		}
		return response.Response{
			Status: false,
			Error:  "Token cannot parsing",
		}
	}

	// Cek apakah token valid
	if !token.Valid {
		if helper.Debug() {
			fmt.Println("Token not valid")
		}
		return response.Response{
			Status: false,
			Error:  "Token not valid",
		}
	}

	// Cek masa berlaku token
	claims, ok := token.Claims.(jwt.MapClaims)
	var expirationTime time.Time
	var iatTime time.Time

	if ok && token.Valid {
		expirationTime = time.Unix(int64(claims["exp"].(float64)), 0)
		iatTime = time.Unix(int64(claims["iat"].(float64)), 0)
	}

	if !iatTime.Before(expirationTime) {
		if helper.Debug() {
			fmt.Println("Error Token valid until:", expirationTime)
		}
		return response.Response{
			Status: false,
			Error:  "Token expired : " + expirationTime.String() + " " + iatTime.String(),
		}
	}

	return response.Response{
		Status:  true,
		Message: "Token valid until:" + expirationTime.String(),
		Data:    claims,
	}

}
