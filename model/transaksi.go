package model

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	NamaPelanggan  string
	NamaProduk     string
	Qty            int
	TotalTransaksi int
	PembuatNota    string
}
