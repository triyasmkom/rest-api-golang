package helpers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	model "server-golang/models/database"
)

var DB *gorm.DB

func Init() {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error init database! ")
	}

	Migration()
}

func Migration() {
	err := DB.AutoMigrate(
		model.User{},
		model.Profile{},
		model.Address{})
	if err != nil {
		panic("Error migrate database!")
	}
}
