package transaksi

import (
	"fmt"
	"strings"
	"tokoku/model"

	"gorm.io/gorm"
)

type TransactionSystem struct {
	DB *gorm.DB
}

func (ts *TransactionSystem) AddTransaction(userName string) (model.Transaksi, bool) {
	var transaction model.Transaksi

	var customerID uint
	fmt.Print("Masukkan ID Pelanggan: ")
	fmt.Scanln(&customerID)

	var customer model.Customer
	if err := ts.DB.First(&customer, customerID).Error; err != nil {
		fmt.Println("Pelanggan tidak ditemukan.")
		return model.Transaksi{}, false
	}

	transaction.NamaPelanggan = customer.Nama

	productQuantities := make(map[uint]int)

	fmt.Println("Masukkan produk dalam transaksi (0 untuk selesai):")
	for {
		var productID uint
		var quantity int

		fmt.Print("Masukkan ID Produk: ")
		fmt.Scanln(&productID)

		if productID == 0 {
			break
		}

		var product model.Produk
		if err := ts.DB.First(&product, productID).Error; err != nil {
			fmt.Println("Produk tidak ditemukan.")
			continue
		}

		fmt.Print("Masukkan Jumlah: ")
		fmt.Scanln(&quantity)

		productQuantities[productID] = quantity

		if product.Stok < quantity {
			fmt.Println("Stok produk tidak mencukupi.")
			continue
		}
		product.Stok -= quantity
		ts.DB.Save(&product)
	}

	total := 0
	qtyItems := []string{}
	for productID, quantity := range productQuantities {
		var product model.Produk
		if err := ts.DB.First(&product, productID).Error; err != nil {
			fmt.Println("Produk tidak ditemukan.")
			return model.Transaksi{}, false
		}
		total += product.HargaProduk * quantity
		qtyItems = append(qtyItems, fmt.Sprintf("%s: %d", product.NamaProduk, quantity))
	}
	transaction.TotalTransaksi = total
	transaction.Qty = strings.Join(qtyItems, ", ")

	transaction.PembuatNota = userName

	if err := ts.DB.Create(&transaction).Error; err != nil {
		fmt.Println("Gagal menyimpan transaksi:", err)
		return model.Transaksi{}, false
	}

	fmt.Println("Transaksi berhasil disimpan.")
	return transaction, true
}

func (vs *TransactionSystem) ViewAllTransaction() ([]model.Transaksi, error) {
	var transactions []model.Transaksi

	err := vs.DB.Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (vs *TransactionSystem) DeleteTransaction() (model.Transaksi, bool) {
	var newTransaksi = new(model.Transaksi)
	fmt.Print("hapus Transaksi dengan ID: ")
	fmt.Scanln(&newTransaksi.ID)

	err := vs.DB.Where("id = ?", newTransaksi.ID).Delete(&newTransaksi).Error
	if err != nil {
		fmt.Println("delete error:")
		return model.Transaksi{}, false
	}
	return model.Transaksi{}, true
}
