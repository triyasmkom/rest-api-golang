package database

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Id          uint   `json:"id" gorm:"primaryKey autoIncrement"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	UserId      uint
}
