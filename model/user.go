package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string
	Alamat   string
	Umur     int
	Gender   string
	Username string
	Password string
	Role     string
}
