package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server-golang/models/response"
)

func Register(context echo.Context) error {

	return context.JSON(http.StatusCreated, response.Response{
		Status:  true,
		Message: "Register success",
		Data:    response.Token{},
	})
}
func Login(context echo.Context) error {
	return context.JSON(http.StatusCreated, response.Response{
		Status:  true,
		Message: "Login success",
		Data:    response.Token{},
	})
}

func Logout(context echo.Context) error {

	return context.JSON(http.StatusCreated, response.Response{
		Status:  true,
		Message: "Logout success",
		Data:    nil,
	})
}
