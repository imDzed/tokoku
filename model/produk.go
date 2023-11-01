package model

import "gorm.io/gorm"

type Produk struct {
	gorm.Model
	NamaProduk  string
	HargaProduk float64 `gorm:"type:decimal(10,2);not null"`
	Deskripsi   string
	Stok        int
	Nama        string
}
