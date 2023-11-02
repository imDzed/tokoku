package model

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	NamaPelanggan  string
	Qty            string
	TotalTransaksi int
	PembuatNota    string
}
