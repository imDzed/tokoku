package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        int       `gorm:"primary_key"`
	Nama      string    `gorm:"type:varchar(255)"`
	Alamat    string    `gorm:"type:text"`
	Gender    string    `gorm:"type:varchar(10)"`
	Username  string    `gorm:"type:varchar(255)"`
	Password  string    `gorm:"type:varchar(255)"`
	IsAdmin   bool      `gorm:"default:false:omitempty"`
	CreatedAt time.Time `gorm:"auto_now_add;type:datetime"`
	UpdatedAt time.Time `gorm:"auto_now;type:datetime"`
}
