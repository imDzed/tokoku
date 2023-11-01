package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Nama   string
	Hp     string
	Alamat string
}