package database

import (
	"gorm.io/gorm"
	"time"
)

type Address struct {
	Id        uint            `json:"id,omitempty" gorm:"primaryKey autoIncrement"`
	Alamat    string          `json:"alamat,omitempty"`
	Kelurahan string          `json:"kelurahan,omitempty"`
	Kecamatan string          `json:"kecamatan,omitempty"`
	Kabupaten string          `json:"kabupaten,omitempty"`
	Provinsi  string          `json:"provinsi,omitempty"`
	UserId    uint            `json:"userId,omitempty" gorm:"unique"`
	CreatedAt *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt *time.Time      `json:"updatedAt,omitempty"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
