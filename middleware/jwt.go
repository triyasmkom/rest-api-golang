package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	mysql "server-golang/configs"
	helper "server-golang/helpers"
	"server-golang/models/database"
	model "server-golang/models/database"
	"server-golang/models/response"
	"strconv"
	"strings"
	"time"
)

var secretKey = []byte(os.Getenv("JWT_KEY"))

func GenerateJwt(data database.User) (response.Data, error) {

	jwtExp := os.Getenv("JWT_EXP")
	expTime, errExp := strconv.Atoi(jwtExp)
	if errExp != nil {
		if helper.Debug() {
			fmt.Println("Error Generate Jwt: ", errExp)
		}
		expTime = 5 // exp 5 hour if env null
	}

	roles := []string{}
	for _, value := range data.Roles {
		roles = append(roles, value.Name)
	}

	iat := time.Now()
	exp := time.Now().Add(time.Hour * time.Duration(expTime))
	claims := jwt.MapClaims{
		"role":     roles,
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
		return response.Data{}, err
	}
	return response.Data{
		Token:   result,
		Created: iat.String(),
		Expired: exp.String(),
	}, nil
}

func VerifyJwt(tokenString string) response.Response {

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

func VerifyUser(data string) response.Response {
	// get user by email

	var getUser model.User
	getUserByEmail := mysql.DB.Where("email=?", data).First(&getUser)

	if getUserByEmail.Error != nil {
		if helper.Debug() {
			fmt.Println(getUserByEmail.Error)
		}

		return response.Response{
			Status: false,
			Error:  "User not found",
		}
	}

	return response.Response{
		Status: true,
		Data:   nil,
	}
}

func JwtAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Validasi Header
		getHeader := c.Request().Header.Get("Authorization")
		token := strings.Split(getHeader, " ")
		if getHeader == "" && token[0] != "Bearer" {
			if helper.Debug() {
				fmt.Println("Error Jwt Auth Header: ", getHeader)
			}
			return c.JSON(http.StatusBadRequest, response.Response{
				Status: false,
				Error:  "Unauthorized",
			})
		}

		// Validasi Jwt
		verifyJwt := VerifyJwt(token[1])
		if !verifyJwt.Status {
			if helper.Debug() {
				fmt.Println("Error Jwt Auth:  ", token[1])
			}
			return c.JSON(http.StatusBadRequest, response.Response{
				Status: false,
				Error:  "Unauthorized",
			})
		}

		// Add data users ke context
		c.Set("users", verifyJwt.Data)
		var email string

		ctxUser := c.Get("users")
		//var getUser response.Claims
		if user, ok := ctxUser.(jwt.MapClaims); ok {
			email = user["email"].(string)
		}

		// Verify Users
		verifyUser := VerifyUser(email)
		fmt.Println(verifyUser.Status, email)
		if !verifyUser.Status {
			if helper.Debug() {
				fmt.Println("Error Jwt Auth:  User not found")
			}
			return c.JSON(http.StatusBadRequest, response.Response{
				Status: false,
				Error:  "User not found",
			})
		}

		// Memanggil handler berikutnya dalam rantai middleware
		if err := next(c); err != nil {
			if helper.Debug() {
				fmt.Println("Error Jwt Auth:  ", err)
			}
			return c.JSON(http.StatusBadRequest, response.Response{
				Status: false,
				Error:  "Unauthorized",
			})
		}
		return nil
	}
}
