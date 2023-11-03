package main

import (
	"fmt"
	"tokoku/auth"
	"tokoku/config"
	"tokoku/customer"
	"tokoku/model"
	"tokoku/product"
	"tokoku/transaksi"
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
	db.AutoMigrate(&model.Transaksi{})

	var auth = auth.AuthSystem{DB: db}
	var product = product.ProductSystem{DB: db}
	var customer = customer.CustomerSystem{DB: db}
	var transaksi = transaksi.TransactionSystem{DB: db}

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
				if result.Role == "pegawai" {
					for permit {
						fmt.Printf("\n\n\t==============\t\n")
						fmt.Printf("\t Menu Pegawai\t\n")
						fmt.Printf("\t==============\t\n\n")
						fmt.Println("1. Tambah Produk")
						fmt.Println("2. Lihat Produk")
						fmt.Println("3. Edit Info Produk")
						fmt.Println("4. Update Stok Produk")
						fmt.Println("5. Tambah Daftar Customer")
						fmt.Println("6. Lihat Daftar Customer")
						fmt.Println("7. Buat Nota Transaksi")
						fmt.Println("8. Lihat Nota Transaksi")
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
									fmt.Printf("\nID Produk: %d\nNama Produk: %s\nHarga: %d\nDeskripsi: %s\nStok: %d\nAddBy: %s\n", p.ID, p.NamaProduk, p.HargaProduk, p.Deskripsi, p.Stok, p.Nama)
								}
							}
						case 3:
							_, permit := product.EditProduct()
							if permit {
								fmt.Println("Produk Berhasil Di Update")
							}

						case 4:
							_, permit := product.EditStokProduct()
							if permit {
								fmt.Println("Stok Produk Berhasil Di Update")
							}
						case 5:
							result, permit := customer.AddCustomer()
							if permit {
								fmt.Printf("\nCustomer %s Telah Berhasil Ditambahkan", result.Nama)
							}
						case 6:
							customers, err := customer.GetCustomers()
							if err != nil {
								fmt.Println("Something Wrong", err)
							} else {
								fmt.Printf("\tDaftar Customer\t\n")
								for _, c := range customers {
									fmt.Printf("\nID Customer: %d\nNama User: %s\nNomor Hp User: %s\nAlamat User: %s\n", c.ID, c.Nama, c.Hp, c.Alamat)
								}
							}
						case 7:
							transaksi, permit := transaksi.AddTransaction(result.Nama)
							if permit {
								fmt.Println(transaksi)
							}
						case 8:
							nota, err := transaksi.ViewAllTransaction()
							if err != nil {
								fmt.Println(nota)
							} else {
								fmt.Printf("\tDaftar Nota\t\n")
								for _, n := range nota {
									fmt.Printf("\nID Nota: %d\nNama Pelanggan: %s\nNama Produk: %s\nQty : %d\nTotal Transaksi: %d\nPembuat Nota: %s\nTanggal Dibuat: %s\n\n",n.ID, n.NamaPelanggan, n.NamaProduk, n.Qty, n.TotalTransaksi, n.PembuatNota, n.CreatedAt.Format("2 January 2006, 15:04"))
								}
							}
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
						fmt.Println("4. Tambah Produk")
						fmt.Println("5. Lihat Produk")
						fmt.Println("6. Edit Info Produk")
						fmt.Println("7. Update Stok Produk")
						fmt.Println("8. Hapus Produk")
						fmt.Println("9. Tambah Daftar Customer")
						fmt.Println("10. Lihat Daftar Customer")
						fmt.Println("11. Hapus Customer")
						fmt.Println("12. Buat Nota Transaksi")
						fmt.Println("13. Lihat Nota Transaksi")
						fmt.Println("14. Hapus Nota Transaksi")
						fmt.Println("15. Update Nota Transaksi")
						fmt.Println("0. Logout")
						fmt.Println("99. Exit")
						fmt.Print("Masukkan Pilihan : ")
						fmt.Scanln(&inputMenu)
						switch inputMenu {
						case 1:
							pegawai, permit := auth.Register()
							if permit {
								fmt.Printf("Pegawai Dengan Nama %s Berhasil Ditambah", pegawai.Nama)
							}
						case 2:
							employees, err := auth.ViewPegawai()
							if err != nil {
								fmt.Println("Gagal mengambil data pegawai:", err.Error())
							} else {
								fmt.Printf("\n\tDaftar Pegawai\t\n")
								for _, employee := range employees {
									fmt.Printf("\nID: %d\nNama: %s\nAlamat: %s\nUmur: %d\nUsername: %s\nPassword: %s\nRole: %s\n",
										employee.ID, employee.Nama, employee.Alamat, employee.Umur, employee.Username, employee.Password, employee.Role)
								}
							}

						case 3:
							_, permit := auth.DeletePegawai()
							if permit {
								fmt.Printf("\nBerhasil Menghapus Pegawai\n")
							}
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
								fmt.Printf("\n\tDaftar Produk\t\n")
								for _, p := range products {
									fmt.Printf("\nID Produk: %d\nNama Produk: %s\nHarga: %d\nDeskripsi: %s\nStok: %d\nAddBy: %s\n", p.ID, p.NamaProduk, p.HargaProduk, p.Deskripsi, p.Stok, p.Nama)
								}
							}
						case 6:
							_, permit := product.EditProduct()
							if permit {
								fmt.Println("Produk Berhasil Di Update")
							}

						case 7:
							_, permit := product.EditStokProduct()
							if permit {
								fmt.Println("Stok Produk Berhasil Di Update")
							}

						case 8:
							_, permit := product.DeleteProduk()
							if permit {
								fmt.Printf("\nBerhasil Menghapus Produk\n")
							}
						case 9:
							result, permit := customer.AddCustomer()
							if permit {
								fmt.Printf("\nCustomer %s Telah Berhasil Ditambahkan", result.Nama)
							}
						case 10:
							customers, err := customer.GetCustomers()
							if err != nil {
								fmt.Println("Something Wrong", err)
							} else {
								fmt.Printf("\n\tDaftar Customer\t\n")
								for _, c := range customers {
									fmt.Printf("\nId User:\t%d\nNama User:\t%s\nNomor Hp User:\t%s\nAlamat User:\t%s\n", c.ID, c.Nama, c.Hp, c.Alamat)
								}
							}
						case 11:
							_, permit := customer.DeleteCustomer()
							if permit {
								fmt.Printf("\n\tCustomer Berhasil Dihapus\t\n")
							}
						case 12:
							transaksi, permit := transaksi.AddTransaction(result.Nama)
							if permit {
								fmt.Println(transaksi)
							}
						case 13:
							nota, err := transaksi.ViewAllTransaction()
							if err != nil {
								fmt.Println(nota)
							} else {
								fmt.Printf("\n\tDaftar Nota\t\n")
								for _, n := range nota {
									fmt.Printf("\nNo Nota: %d\nNama Pelanggan: %s\nQty Detail: %d\nTotal Transaksi: %d\nPembuat Nota: %s\nTanggal Dibuat: %s\n\n", n.ID, n.NamaPelanggan, n.Qty, n.TotalTransaksi, n.PembuatNota, n.CreatedAt.Format("2 January 2006, 15:04"))
								}
							}
						case 14:
							_, permit := transaksi.DeleteTransaction()
							if permit {
								fmt.Printf("\nBerhasil Menghapus Transaksi\n")
							}
						case 15:
							_, permit := transaksi.UpdateTransaction()
							if permit {
								fmt.Println("Transaksi Berhasil Di Update")
							}
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
