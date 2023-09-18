package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
	s_user "server-golang/services"
)

func AddProfile(context echo.Context) error {
	add := s_user.AddProfile(context)
	if !add.Status {
		return context.JSON(http.StatusBadRequest, add)
	}
	return context.JSON(http.StatusCreated, add)
}

func UpdateProfile(context echo.Context) error {
	update := s_user.UpdateProfile(context)
	if !update.Status {
		return context.JSON(http.StatusBadRequest, update)
	}
	return context.JSON(http.StatusOK, update)
}

func AddAddress(context echo.Context) error {
	add := s_user.AddAddress(context)
	if !add.Status {
		return context.JSON(http.StatusBadRequest, add)
	}
	return context.JSON(http.StatusCreated, add)
}

func UpdateAddress(context echo.Context) error {
	update := s_user.UpdateAddress(context)

	if !update.Status {
		return context.JSON(http.StatusBadRequest, update)
	}

	return context.JSON(http.StatusOK, update)
}

func GetUser(context echo.Context) error {
	update := s_user.GetUser(context)

	if !update.Status {
		return context.JSON(http.StatusBadRequest, update)
	}

	return context.JSON(http.StatusOK, update)
}
