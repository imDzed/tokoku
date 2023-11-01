package model

import "gorm.io/gorm"

type Produk struct {
	gorm.Model
	Nama      string
	Harga     string
	Deskripsi string
	Stok      string
	AddBy     string
}
