package services

import (
	"fmt"
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"reflect"
	mysql "server-golang/configs"
	helper "server-golang/helpers"
	model "server-golang/models/database"
	"server-golang/models/response"
)

func AddProfile(context echo.Context) response.Response {
	body := new(model.Profile)
	context.Bind(body)

	//var getUser response.Claims
	var email string
	ctxUser := context.Get("users")
	if user, ok := ctxUser.(jwt2.MapClaims); ok {
		email = user["email"].(string)
	}

	// get user by email
	var getUser model.User
	getUserByEmail := mysql.DB.Where("email = ?", email).First(&getUser)
	if getUserByEmail.Error != nil {
		if helper.Debug() {
			fmt.Println(getUserByEmail.Error)
		}

		return response.Response{
			Status: false,
			Error:  "User not found",
		}
	}

	body.UserId = getUser.Id
	save := mysql.DB.Create(&body)
	if save.Error != nil {
		if helper.Debug() {
			fmt.Println("Error Add Profile: ", save.Error)
		}

		if gorm.ErrDuplicatedKey != nil {
			return response.Response{
				Status: false,
				Error:  "Cannot add profile again, please your edit this",
			}
		}

		return response.Response{
			Status: false,
			Error:  "Error Add Profile",
		}
	}

	return response.Response{
		Status:  true,
		Message: "Add Profile User Success",
		Data: response.Data{
			FirstName:   body.FirstName,
			LastName:    body.LastName,
			PhoneNumber: body.PhoneNumber,
		},
	}
}

func UpdateProfile(context echo.Context) response.Response {
	body := new(model.Profile)
	context.Bind(body)

	//var getUser response.Claims
	var email string
	ctxUser := context.Get("users")
	if user, ok := ctxUser.(jwt2.MapClaims); ok {
		email = user["email"].(string)
	}

	// get user by email
	var getUser model.User
	getUserByEmail := mysql.DB.Where("email = ?", email).First(&getUser)
	if getUserByEmail.Error != nil {
		if helper.Debug() {
			fmt.Println(getUserByEmail.Error)
		}

		return response.Response{
			Status: false,
			Error:  "User not found",
		}
	}

	// update profile
	update := mysql.DB.Where("id=?", getUser.Profile.Id).Updates(&body)
	if update.Error != nil {
		if helper.Debug() {
			fmt.Println("Error Update Profile: ", reflect.TypeOf(update.Error))
		}

		return response.Response{
			Status: false,
			Error:  "Error Update Profile",
		}
	}

	return response.Response{
		Status:  true,
		Message: "Update Profile User Success",
		Data: response.Data{
			FirstName:   body.FirstName,
			LastName:    body.LastName,
			PhoneNumber: body.PhoneNumber,
		},
	}
}

func AddAddress(context echo.Context) response.Response {
	body := new(model.Address)
	context.Bind(body)

	//var getUser response.Claims
	var email string
	ctxUser := context.Get("users")
	if user, ok := ctxUser.(jwt2.MapClaims); ok {
		email = user["email"].(string)
	}

	// get user by email
	var getUser model.User
	getUserByEmail := mysql.DB.Where("email = ?", email).First(&getUser)
	if getUserByEmail.Error != nil {
		if helper.Debug() {
			fmt.Println(getUserByEmail.Error)
		}

		return response.Response{
			Status: false,
			Error:  "User not found",
		}
	}

	body.UserId = getUser.Id
	save := mysql.DB.Create(&body)
	if save.Error != nil {
		if helper.Debug() {
			fmt.Println("Error Add Address: ", save.Error)
		}

		if gorm.ErrDuplicatedKey != nil {
			return response.Response{
				Status: false,
				Error:  "Cannot add address again, please your edit this",
			}
		}

		return response.Response{
			Status: false,
			Error:  "Error Add Address",
		}
	}

	return response.Response{
		Status:  true,
		Message: "Add Address User Success",
		Data: response.Data{
			Alamat:    body.Alamat,
			Kelurahan: body.Kelurahan,
			Kecamatan: body.Kecamatan,
			Kabupaten: body.Kabupaten,
			Provinsi:  body.Provinsi,
		},
	}
}

func UpdateAddress(context echo.Context) response.Response {
	body := new(model.Address)
	context.Bind(body)

	//var getUser response.Claims
	var email string
	ctxUser := context.Get("users")
	if user, ok := ctxUser.(jwt2.MapClaims); ok {
		email = user["email"].(string)
	}

	// get user by email
	var getUser model.User
	getUserByEmail := mysql.DB.Where("email = ?", email).First(&getUser)
	if getUserByEmail.Error != nil {
		if helper.Debug() {
			fmt.Println(getUserByEmail.Error)
		}

		return response.Response{
			Status: false,
			Error:  "User not found",
		}
	}

	// update
	update := mysql.DB.Where("id=?", getUser.Address.Id).Updates(&body)
	if update.Error != nil {
		if helper.Debug() {
			fmt.Println("Error Update Address: ", update.Error)
		}

		return response.Response{
			Status: false,
			Error:  "Error Update Address",
		}
	}

	return response.Response{
		Status:  true,
		Message: "Update Address User Success",
		Data: response.Data{
			Alamat:    body.Alamat,
			Kelurahan: body.Kelurahan,
			Kecamatan: body.Kecamatan,
			Kabupaten: body.Kabupaten,
			Provinsi:  body.Provinsi,
		},
	}
}
