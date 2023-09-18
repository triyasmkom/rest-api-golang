package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server-golang/models/response"
	s_user "server-golang/services"
)

func Register(context echo.Context) error {

	save := s_user.Register(context)
	if !save.Status {
		return context.JSON(http.StatusBadRequest, save)
	}

	return context.JSON(http.StatusCreated, save)
}
func Login(context echo.Context) error {
	save := s_user.Login(context)

	if !save.Status {
		return context.JSON(http.StatusBadRequest, save)
	}

	return context.JSON(http.StatusOK, save)
}

// not use
func Logout(context echo.Context) error {

	return context.JSON(http.StatusCreated, response.Response{
		Status:  true,
		Message: "Logout success",
		Data:    response.Data{},
	})
}
