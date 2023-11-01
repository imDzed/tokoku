package main

import (
	"fmt"
	"tokoku/auth"
	"tokoku/config"
	"tokoku/model"
)

func main() {
	var inputMenu int
	db, err := config.InitDB()
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Produk{})

	var auth = auth.AuthSystem{DB: db}

	for{
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("99. Exit")
		fmt.Print("Masukkan Pilihan: ")
		fmt.Scanln(&inputMenu)
		switch inputMenu {
		case 1:
			result, permit := auth.Login()
			if permit{
				fmt.Println("selamat datang", result.Nama)
				var pilihmenu int 
				fmt.Println("1. Tambah Barang")
				fmt.Println("2. edit info barang")
				fmt.Println("3. update stok barang")
				fmt.Println("4. tambah daftar customer")
				fmt.Println("5. buat nota transaksi")
				fmt.Print("Masukkan Pilihan : ")
				fmt.Scanln(&pilihmenu)
				switch pilihmenu {
				case 1:
				case 2:
				case 3:
				case 4:
				case 5:
					return
				}
			}
		case 2:
			result, permit := auth.Register()
			if permit {
				fmt.Println(result)
			}
		case 99:
			fmt.Println("Thank you ....")
			return
		default:
		}
	}

}