package model

import "gorm.io/gorm"

type Produk struct {
	gorm.Model
	NamaProduk  string
	HargaProduk int
	Deskripsi   string
	Stok        int
	Nama        string
}
