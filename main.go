package main

import (
	"fmt"
	"tokoku/auth"
	"tokoku/config"
	"tokoku/customer"
	"tokoku/model"
	"tokoku/product"
)

func main() {
	var inputMenu int
	db, err := config.InitDB()
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Produk{})
	db.AutoMigrate(&model.Customer{})

	var auth = auth.AuthSystem{DB: db}
	var product = product.ProductSystem{DB: db}
	var customer = customer.CustomerSystem{DB: db}

	for {
		fmt.Println("1. Login")
		fmt.Println("99. Exit")
		fmt.Print("Masukkan Pilihan: ")
		fmt.Scanln(&inputMenu)
		switch inputMenu {
		case 1:
			result, permit := auth.Login()
			if permit {
				fmt.Println("Selamat Datang ", result.Nama)
				if result.Role == "" {
					for permit {
						fmt.Printf("\n\n\t==============\t\n")
						fmt.Printf("\t Menu Pegawai\t\n")
						fmt.Printf("\t==============\t\n\n")
						fmt.Println("1. Tambah Barang")
						fmt.Println("2. Lihat Barang")
						fmt.Println("3. Edit Info Barang")
						fmt.Println("4. Update Stok Barang")
						fmt.Println("5. Tambah Daftar Customer")
						fmt.Println("6. Buat Nota Transaksi")
						fmt.Println("0. Logout")
						fmt.Println("99. Exit")
						fmt.Print("Masukkan Pilihan : ")
						fmt.Scanln(&inputMenu)
						switch inputMenu {
						case 1:
							result, permit := product.AddProduct(result.Nama)
							if permit {
								fmt.Printf("\n%s Telah Berhasil Ditambahkan !!\n", result.NamaProduk)
							}
						case 2:
							products, err := product.ViewAllProducts()
							if err != nil {
								fmt.Println("Gagal mengambil produk:", err)
							} else {
								fmt.Println("Daftar Produk:")
								for _, p := range products {
									fmt.Printf("\nNama Produk: %s\nHarga: %.2f\nDeskripsi: %s\nStok: %d\nAddBy: %s\n\n", p.NamaProduk, p.HargaProduk, p.Deskripsi, p.Stok, p.Nama)
								}
							}
						case 3:
						case 4:
						case 5:
							result, permit := customer.AddCustomer()
							if permit {
								fmt.Printf(result.Nama, "\n%s Telah Berhasil Ditambahkan !!\n")
							}
						case 6:
						case 0:
							permit = false
						case 99:
							fmt.Println("Thank you ....")
							return
						}
					}
				} else if result.Role == "admin" {
					for permit {
						fmt.Printf("\n\n\t=============\t\n")
						fmt.Printf("\t Menu Admin\t\n")
						fmt.Printf("\t=============\t\n\n")
						fmt.Println("1. Tambah Pegawai")
						fmt.Println("2. Lihat Pegawai")
						fmt.Println("3. Hapus Pegawai")
						fmt.Println("4. Tambah Barang")
						fmt.Println("5. Lihat Barang")
						fmt.Println("6. Edit Info Barang")
						fmt.Println("7. Update Stok Barang")
						fmt.Println("8. Hapus Barang")
						fmt.Println("9. Tambah Daftar Customer")
						fmt.Println("10. Lihat Daftar Customer")
						fmt.Println("11. Hapus Customer")
						fmt.Println("12. Buat Nota Transaksi")
						fmt.Println("13. Lihat Nota Transaksi")
						fmt.Println("14. Hapus Nota Transaksi")
						fmt.Println("0. Logout")
						fmt.Println("99. Exit")
						fmt.Print("Masukkan Pilihan : ")
						fmt.Scanln(&inputMenu)
						switch inputMenu {
						case 1:
							result, permit := auth.Register()
							if permit {
								fmt.Println(result)
							}
						case 2:
						case 3:
						case 4:
							result, permit := product.AddProduct(result.Nama)
							if permit {
								fmt.Printf("\n%s Telah Berhasil Ditambahkan !!\n", result.NamaProduk)
							}
						case 5:
							products, err := product.ViewAllProducts()
							if err != nil {
								fmt.Println("Gagal mengambil produk:", err)
							} else {
								fmt.Println("Daftar Produk:")
								for _, p := range products {
									fmt.Printf("\nNama Produk: %s\nHarga: %.2f\nDeskripsi: %s\nStok: %d\nAddBy: %s\n\n", p.NamaProduk, p.HargaProduk, p.Deskripsi, p.Stok, p.Nama)
								}
							}
						case 6:
						case 7:
						case 8:
						case 9:
							_, permit := customer.AddCustomer()
							if permit {
								fmt.Printf("\n Customer Berhasil Ditambahkan !!\n")
							}
							fmt.Println("Kembali ke menu sebelumnya? (y/n)")
							var back string
							fmt.Scanln(&back)

							if back == "y" {
								inputMenu = 0
							}
						case 10:
							customers, err := customer.GetCustomers()

							if err != nil {
								fmt.Println("Gagal mengambil customer:", err)
							} else {
								fmt.Println("\nDaftar Customer:")
								for _, c := range customers {
									fmt.Printf("\nNama Customer: %s\nNomor HP: %s\nAlamat: %s\n\n", c.Nama, c.Hp, c.Alamat)
								}
							}
							fmt.Println("Kembali ke menu sebelumnya? (y/n)")
							var back string
							fmt.Scanln(&back)

							if back == "y" {
								inputMenu = 0
							}
						case 11:
						case 12:
						case 13:
						case 14:
						case 0:
							permit = false
						case 99:
							fmt.Println("Thank you ....")
							return
						}
					}
				}
			}

		case 99:
			fmt.Println("Thank you ....")
			return
		default:
		}
	}

}
