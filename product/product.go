package product

import (
	"bufio"
	"fmt"
	"os"
	"tokoku/model"

	"gorm.io/gorm"
)

type ProductSystem struct {
	DB *gorm.DB
}

func (ps *ProductSystem) AddProduct(userName string) (model.Produk, bool) {
	var newProduct = new(model.Produk)
	fmt.Println("Masukkan Nama Produk")
	fmt.Scanln(&newProduct.NamaProduk)
	fmt.Println("Masukkan Harga Produk")
	fmt.Scanln(&newProduct.HargaProduk)
	fmt.Println("Masukkan Deskripsi Produk")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	newProduct.Deskripsi = scanner.Text()
	fmt.Println("Masukkan Stok Produk")
	fmt.Scanln(&newProduct.Stok)

	newProduct.Nama = userName

	err := ps.DB.Create(newProduct).Error
	if err != nil {
		fmt.Println("Input Produk Gagal:", err.Error())
		return model.Produk{}, false
	}

	return *newProduct, true
}
