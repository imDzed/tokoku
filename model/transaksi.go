package model

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	NamaProduk     string
	NamaPelanggan  string
	Qty            string
	TotalTransaksi string
	PembuatNota    string
}
