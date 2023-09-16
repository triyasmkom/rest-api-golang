package database

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Id   uint   `json:"id" gorm:"primaryKey autoIncrement"`
	Name string `json:"name" gorm:"unique"`
	User []User `gorm:"many2many:user_role"`
}
