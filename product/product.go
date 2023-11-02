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
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan Nama Produk: ")
	scanner.Scan()
	newProduct.NamaProduk = scanner.Text()

	fmt.Print("Masukkan Harga Produk: ")
	fmt.Scanln(&newProduct.HargaProduk)

	fmt.Print("Masukkan Deskripsi Produk: ")
	scanner.Scan()
	newProduct.Deskripsi = scanner.Text()

	fmt.Print("Masukkan Stok Produk: ")
	fmt.Scanln(&newProduct.Stok)

	newProduct.Nama = userName

	err := ps.DB.Create(newProduct).Error
	if err != nil {
		fmt.Println("Input Produk Gagal:", err.Error())
		return model.Produk{}, false
	}

	return *newProduct, true
}

func (ps *ProductSystem) ViewAllProducts() ([]model.Produk, error) {
	var products []model.Produk

	err := ps.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}


func (dp *ProductSystem) DeleteProduk() (model.Produk, bool) {
	var newProduk = new(model.Produk)
	fmt.Print("Hapus User Dengan ID: ")
	fmt.Scanln(&newProduk.ID)

	err := dp.DB.Where("id = ?", newProduk.ID).Delete(&newProduk).Error
	if err != nil {
		fmt.Println("delete error:")
		return model.Produk{}, false
	}
	return model.Produk{}, true
}