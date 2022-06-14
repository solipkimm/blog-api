package main

import (
	"config"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	config.Load()
	Connect()
}

func Connect() {
	db, err := gorm.Open(config.DB_DRIVER, config.DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Successfully Connected to MySQL database")
}
