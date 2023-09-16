package database

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint            `json:"id" gorm:"primaryKey autoIncrement"`
	Username  string          `json:"username,omitempty"`
	Email     string          `json:"email,omitempty" gorm:"unique"`
	Password  string          `json:"password,omitempty"`
	IsVerify  bool            `json:"isVerify,omitempty" gorm:"default:false"`
	CreatedAt *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt *time.Time      `json:"updatedAt,omitempty"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	Roles     []Role          `gorm:"many2many:user_role" json:"roles,omitempty"`
	Profile   Profile         `gorm:"foreignKey:UserId; references:id" json:"profile,omitempty"`
	Address   Address         `gorm:"foreignKey:UserId; references:id" json:"address,omitempty"`
}
