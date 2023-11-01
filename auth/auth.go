package auth

import (
	"fmt"
	"tokoku/model"

	"gorm.io/gorm"
)

type AuthSystem struct {
	DB *gorm.DB
}

func (as *AuthSystem) Register() (model.User, bool) {
	var newUser = new(model.User)
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&newUser.Nama)
	fmt.Print("Masukkan Alamat: ")
	fmt.Scanln(&newUser.Alamat)
	fmt.Print("masukkan Umur: ")
	fmt.Scanln(&newUser.Umur)
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&newUser.Username)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&newUser.Password)
	fmt.Println("Masukkan Role: ")
	fmt.Scanln(&newUser.Role)
	err := as.DB.Create(newUser).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.User{}, false
	}
	return *newUser, true
}

func (as *AuthSystem) Login() (model.User, bool) {

	var currentUser = new(model.User)

	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&currentUser.Username)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&currentUser.Password)

	qry := as.DB.Where("username = ? AND password = ?", currentUser.Username, currentUser.Password).Take(currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("login process error:", err.Error())
		return model.User{}, false
	}

	return *currentUser, true
}
