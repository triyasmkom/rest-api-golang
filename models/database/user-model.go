package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint    `json:"id" gorm:"primaryKey autoIncrement"`
	Username string  `json:"username"`
	Email    string  `json:"email" gorm:"unique"`
	Password string  `json:"password"`
	IsVerify bool    `json:"isVerify" gorm:"default:false"`
	Roles    []Role  `gorm:"many2many:user_role"`
	Profile  Profile `gorm:"foreignKey:UserId; references:id"`
	Address  Address `gorm:"foreignKey:UserId; references:id"`
}
