package database

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	Id        uint            `json:"id,omitempty" gorm:"primaryKey autoIncrement"`
	Name      string          `json:"name,omitempty" gorm:"unique"`
	User      []User          `json:"user,omitempty" gorm:"many2many:user_role"`
	CreatedAt *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt *time.Time      `json:"updatedAt,omitempty"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
