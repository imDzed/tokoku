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
			return transaction, false
		}

		fmt.Print("Masukkan Jumlah: ")
		fmt.Scanln(&quantity)

		productQuantities[productID] = quantity

		if product.Stok < quantity {
			fmt.Println("Stok produk tidak mencukupi.")
			return transaction, false
		}
		product.Stok -= quantity
		ts.DB.Save(&product)
	}

	total := 0
	totalqty := 0
	qtyItems := []string{}
	for productID, quantity := range productQuantities {
		var product model.Produk
		if err := ts.DB.First(&product, productID).Error; err != nil {
			fmt.Println("Produk tidak ditemukan.")
			return model.Transaksi{}, false
		}
		total += product.HargaProduk * quantity
		qtyItems = append(qtyItems, fmt.Sprintf("\n\t\t%s: %d", product.NamaProduk, quantity))
		totalqty += quantity
	}
	transaction.TotalTransaksi = total
	transaction.NamaProduk = strings.Join(qtyItems, "")
	transaction.Qty = totalqty

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

func (ts *TransactionSystem) UpdateTransaction() (model.Transaksi, bool) {
	var updatedTransaksi model.Transaksi

	fmt.Print("Masukkan ID Transaksi yang akan diperbarui: ")
	var transaksiID uint
	fmt.Scanln(&transaksiID)

	err := ts.DB.Where("id = ?", transaksiID).First(&updatedTransaksi).Error
	if err != nil {
		fmt.Println("Transaksi tidak ditemukan:", err.Error())
		return model.Transaksi{}, false
	}

	productQuantities := make(map[uint]int)

	fmt.Println("Masukkan produk dalam transaksi (0 untuk selesai):")
	for {
		var productID uint
		var quantity int

		fmt.Print("Masukkan ID Produk (0 untuk selesai): ")
		fmt.Scanln(&productID)

		if productID == 0 {
			break
		}

		var product model.Produk
		if err := ts.DB.First(&product, productID).Error; err != nil {
			fmt.Println("Produk tidak ditemukan.")
			return updatedTransaksi, false
		}

		fmt.Printf("Masukkan Jumlah produk %s : ", product.NamaProduk)
		fmt.Scanln(&quantity)

		if product.Stok < quantity {
			fmt.Println("Stok produk tidak mencukupi.")
			return updatedTransaksi, false
		}

		stokChange := productQuantities[productID] - quantity
		product.Stok += stokChange

		productQuantities[productID] = quantity
		product.Stok -= quantity
		ts.DB.Save(&product)
	}

	total := 0
	totalqty := 0
	qtyItems := []string{}
	for productID, quantity := range productQuantities {
		var product model.Produk
		if err := ts.DB.First(&product, productID).Error; err != nil {
			fmt.Println("Produk tidak ditemukan.")
			return model.Transaksi{}, false
		}
		total += product.HargaProduk * quantity
		qtyItems = append(qtyItems, fmt.Sprintf("\n\t\t%s: %d", product.NamaProduk, quantity))
		totalqty += quantity
	}
	updatedTransaksi.TotalTransaksi = total
	updatedTransaksi.NamaProduk = strings.Join(qtyItems, "")
	updatedTransaksi.Qty = totalqty

	if err := ts.DB.Save(&updatedTransaksi).Error; err != nil {
		fmt.Println("Gagal menyimpan perubahan:", err.Error())
		return model.Transaksi{}, false
	}

	fmt.Println("Transaksi berhasil diperbarui.")
	return updatedTransaksi, true
}
