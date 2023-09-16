package services

import (
	"fmt"
	mysql "server-golang/configs"
	helper "server-golang/helpers"
	jwt "server-golang/middleware"
	model "server-golang/models/database"
	"server-golang/models/request"
	"server-golang/models/response"
)

func Register(body *request.Register) interface{} {
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

		if helper.Debug() {
			fmt.Println(errToken)
		}

		return response.Response{
			Status: false,
			Error:  "Error Generate Jwt",
		}
	}

	// Save user
	save := mysql.DB.Create(&data)
	if save.Error != nil {
		if helper.Debug() {
			fmt.Println(save.Error)
		}
		return response.Response{
			Status: false,
			Error:  "Error Register",
		}
	}

	// Save Role User
	err := mysql.DB.Model(&data).Association("Roles").Append(&role)
	if err != nil {

		if helper.Debug() {
			fmt.Println(err)
		}

		return response.Response{
			Status: false,
			Error:  "Error Register",
		}
	}

	return response.Response{
		Status:  true,
		Message: "Save User Success",
		Data:    token,
	}
}
func Login(body *request.Login) interface{} {
	// get user by email
	var getUser model.User
	mysql.DB.Where("email = ?", body.Email).First(&getUser)

	// Validasi password
	_, err := helper.VerifyPassword(body.Password, getUser.Password)
	if err != nil {

		if helper.Debug() {
			fmt.Println(err)
		}

		return response.Response{
			Status: false,
			Error:  "Wrong Email or Password",
		}
	}

	// Generate Jwt
	token, errToken := jwt.GenerateJwt(getUser)
	if errToken != nil {

		if helper.Debug() {
			fmt.Println(errToken)
		}

		return response.Response{
			Status: false,
			Error:  "Error Generate Jwt",
		}
	}

	return response.Response{
		Status:  true,
		Message: "Login User Success",
		Data:    token,
	}
}
