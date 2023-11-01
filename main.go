package main

import (
	"fmt"
	"tokoku/config"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		fmt.Println(db)
	}

}