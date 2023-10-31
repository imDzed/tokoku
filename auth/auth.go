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
	var addUser = new(model.User)

	fmt.Println("Masukkan nama User: ")
	fmt.Scanln(&addUser.Nama)

	fmt.Println("Masukkan alamat User: ")
	fmt.Scanln(&addUser.Alamat)

	fmt.Println("Masukkan gender User (Laki-laki/Perempuan): ")
	fmt.Scanln(&addUser.Gender)

	fmt.Println("Masukkan username User: ")
	fmt.Scanln(&addUser.Username)

	fmt.Println("Masukkan password User: ")
	fmt.Scanln(&addUser.Password)

	fmt.Println("Status akun (true = 1, false = 0)")
	fmt.Scanln(&addUser.IsAdmin)

	err := as.DB.Create(addUser).Error
	if err != nil {
		fmt.Println("Gagal menambahkan User:", err.Error())
		return model.User{}, false
	}

	return *addUser, true
}
