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
		fmt.Println(db, err.Error())
		return
	}

	db.AutoMigrate(&model.User{})

	var auth = auth.AuthSystem{DB: db}

	for {
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("99. Exit")
		fmt.Print("Masukkan pilihan:")
		fmt.Scanln(&inputMenu)
		switch inputMenu {
		case 2:
			result, permit := auth.Register()
			if permit {
				fmt.Println(result)
			}
		case 99:
			fmt.Println("Thank you....")
			return
		default:

		}
	}

}
