package auth

import (
	"bufio"
	"fmt"
	"os"
	"tokoku/model"

	"gorm.io/gorm"
)

type AuthSystem struct {
	DB *gorm.DB
}

func (ar *AuthSystem) Register() (model.User, bool) {
	var newUser = new(model.User)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan nama: ")
	scanner.Scan()
	newUser.Nama = scanner.Text()

	fmt.Print("Masukkan Alamat: ")
	scanner.Scan()
	newUser.Alamat = scanner.Text()

	fmt.Print("masukkan Umur: ")
	fmt.Scanln(&newUser.Umur)

	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&newUser.Username)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&newUser.Password)

	if newUser.Role == "" {
		newUser.Role = "pegawai"
	}

	err := ar.DB.Create(newUser).Error
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

func (av *AuthSystem) ViewPegawai() ([]model.User, error) {
	var employees []model.User
	err := av.DB.Where("Role = ?", "pegawai").Find(&employees).Error
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (vs *AuthSystem) DeletePegawai() (model.User, bool) {
	var newPegawai = new(model.User)
	fmt.Print("Hapus Pegawai dengan ID: ")
	fmt.Scanln(&newPegawai.ID)

	err := vs.DB.Where("id = ?", newPegawai.ID).Delete(&newPegawai).Error
	if err != nil {
		fmt.Println("delete error:")
		return model.User{}, false
	}
	return model.User{}, true
}
