package model

import "gorm.io/gorm"

type Produk struct {
	gorm.Model
	NamaProduk string
	Harga      string
	Deskripsi  string
	Stok       string
	Nama       string
}
