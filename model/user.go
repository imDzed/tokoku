package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string
	Alamat   string
	Umur     string
	Gender   string
	Username string
	Password string
	Role     string
}
