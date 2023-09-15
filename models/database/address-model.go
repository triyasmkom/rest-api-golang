package database

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Id        uint   `json:"id" gorm:"primaryKey autoIncrement"`
	Alamat    string `json:"alamat"`
	Kelurahan string `json:"kelurahan"`
	Kecamatan string `json:"kecamatan"`
	Kabupaten string `json:"kabupaten"`
	Provinsi  string `json:"provinsi"`
	UserId    uint
}
