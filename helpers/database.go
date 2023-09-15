package helpers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	model "server-golang/models/database"
)

var DB *gorm.DB

func Init() {
	//dsn := fmt.Sprintln(
	//	"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
	//	os.Getenv("DB_USER"),
	//	os.Getenv("DB_PASSWORD"),
	//	os.Getenv("DB_HOST"),
	//	os.Getenv("DB_PORT"),
	//	os.Getenv("DB_NAME"),
	//)

	dsn := "root:12345@tcp(localhost:3306)/rest_api_gin?charset=utf8mb4&parseTime=True&loc=Local"

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
