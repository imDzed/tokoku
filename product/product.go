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




func (ps *ProductSystem) EditProduct() (model.Produk, bool) {
    var editedProduct model.Produk
	scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Masukkan ID Produk yang akan diedit: ")
    var productID uint
    fmt.Scanln(&productID)

    err := ps.DB.Where("id = ?", productID).First(&editedProduct).Error
    if err != nil {
        fmt.Println("Produk tidak ditemukan:", err.Error())
        return model.Produk{}, false
    }

    var userInput string

    fmt.Printf("Nama Produk (sebelumnya: %s): ", editedProduct.NamaProduk)
	if userInput != "" {
        editedProduct.NamaProduk = userInput
    }
	scanner.Scan()
	editedProduct.NamaProduk = scanner.Text()
    

    fmt.Printf("Harga Produk (sebelumnya: %d): ", editedProduct.HargaProduk)
    fmt.Scanln(&userInput)
    if userInput != "" {
        fmt.Sscan(userInput, &editedProduct.HargaProduk)
    }

    fmt.Printf("Deskripsi Produk (sebelumnya: %s): ", editedProduct.Deskripsi)
	if userInput != "" {
        editedProduct.Deskripsi = userInput
    }
	scanner.Scan()
	editedProduct.Deskripsi = scanner.Text()
    

    fmt.Printf("Stok Produk (sebelumnya: %d): ", editedProduct.Stok)
    fmt.Scanln(&userInput)
    if userInput != "" {
        fmt.Sscan(userInput, &editedProduct.Stok)
    }

    err = ps.DB.Save(&editedProduct).Error
    if err != nil {
        fmt.Println("Gagal menyimpan perubahan:", err.Error())
        return model.Produk{}, false
    }

    return editedProduct, true
}


func (ps *ProductSystem) EditStokProduct() (model.Produk, bool) {
	var editStokProduct model.Produk

	fmt.Print("Masukkan ID Produk yang akan diedit: ")
    var productID uint
    fmt.Scanln(&productID)

    err := ps.DB.Where("id = ?", productID).First(&editStokProduct).Error
    if err != nil {
        fmt.Println("Produk tidak ditemukan:", err.Error())
        return model.Produk{}, false
    }

    var userInput string

	fmt.Printf("Stok Produk (sebelumnya: %d): ", editStokProduct.Stok)
    fmt.Scanln(&userInput)
    if userInput != "" {
        fmt.Sscan(userInput, &editStokProduct.Stok)
    }

    err = ps.DB.Save(&editStokProduct).Error
    if err != nil {
        fmt.Println("Gagal menyimpan perubahan:", err.Error())
        return model.Produk{}, false
    }

    return editStokProduct, true

}