package services

import (
	mysql "server-golang/configs"
	helper "server-golang/helpers"
	jwt "server-golang/middleware"
	model "server-golang/models/database"
	"server-golang/models/request"
	"server-golang/models/response"

	"github.com/labstack/echo/v4"
)

func SaveUser(body *request.Register) interface{} {
	data := model.User{
		Email:    body.Email,
		Password: helper.HashPassword(body.Password),
		Username: body.Username,
	}
	role := model.Role{
		Name: "User",
	}

	token, errToken := jwt.GenerateJwt(data)
	if errToken != nil {
		return response.Response{
			Status:  false,
			Message: "Error Generate Jwt",
			Data:    errToken,
		}
	}

	// Save user
	save := mysql.DB.Create(&data)
	if save.Error != nil {
		return response.Response{
			Status:  false,
			Message: "Error Save User",
			Data:    save.Error,
		}
	}

	// Save Role User
	err := mysql.DB.Model(&data).Association("Roles").Append(&role)
	if err != nil {
		return response.Response{
			Status:  false,
			Message: "Error Add Role",
			Data: echo.Map{
				"Error": err.Error(),
			},
		}
	}

	return response.Response{
		Status:  true,
		Message: "Save User Success",
		Data: echo.Map{
			"email":    data.Email,
			"username": data.Username,
			"token":    token,
		},
	}
}
