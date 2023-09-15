package services

import (
	hash "server-golang/helpers"
	mysql "server-golang/helpers"
	userMod "server-golang/models/database"
	"server-golang/models/request"
	"server-golang/models/response"
)

func SaveUser(body *request.Register) interface{} {
	var data userMod.User
	data.Email = body.Email
	data.Password = hash.HashPassword(body.Password)
	data.Username = body.Username

	save := mysql.DB.Create(&data)
	if save.Error != nil {
		return response.Response{
			Status:  false,
			Message: "Error Save User",
			Data:    save.Error,
		}
	}
	return response.Response{
		Status:  true,
		Message: "Save User Success",
		Data: response.Token{
			Token:   "",
			Created: "",
			Expired: "",
		},
	}
}
