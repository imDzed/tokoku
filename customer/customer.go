package customer

import (
	"bufio"
	"fmt"
	"os"
	"tokoku/model"

	"gorm.io/gorm"
)

type CustomerSystem struct {
	DB *gorm.DB
}

func (cs *CustomerSystem) AddCustomer() (model.Customer, bool) {
	var newCustomer = new(model.Customer)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan Nama Customer: ")
	scanner.Scan()
	newCustomer.Nama = scanner.Text()

	fmt.Print("Masukkan Nomor HP: ")
	fmt.Scanln(&newCustomer.Hp)

	fmt.Print("Masukkan Alamat: ")
	scanner.Scan()
	newCustomer.Alamat = scanner.Text()

	err := cs.DB.Create(newCustomer).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.Customer{}, false
	}
	return *newCustomer, true
}

func (cs *CustomerSystem) GetCustomers() ([]model.Customer, error) {
	customers := []model.Customer{}
	err := cs.DB.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (cs *CustomerSystem) DeleteCustomer(id int) ([]model.Customer, error) {
	// Query customer berdasarkan id
	customer := model.Customer{}
	err := cs.DB.Where("id = ?", id).First(&customer).Error
	if err != nil {
		return nil, err
	}

	// Hapus customer
	err = cs.DB.Delete(&customer).Error
	if err != nil {
		return nil, err
	}

	// Kembalikan data customer
	return cs.GetCustomers()
}
