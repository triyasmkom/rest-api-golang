package database

import (
	"gorm.io/gorm"
	"time"
)

type Profile struct {
	Id          uint            `json:"id,omitempty" gorm:"primaryKey autoIncrement"`
	FirstName   string          `json:"firstName,omitempty"`
	LastName    string          `json:"lastName,omitempty"`
	PhoneNumber string          `json:"phoneNumber,omitempty"`
	UserId      uint            `json:"userId,omitempty"`
	CreatedAt   *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time      `json:"updatedAt,omitempty"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
