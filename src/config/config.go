package config

import (
	"fmt"
	"log"
	"os"
)

var (
	DB_DRIVER = ""
	DB_URL    = ""
)

func Load() {
	var err error
	if err != nil {
		log.Fatal(err)
	}

	DB_DRIVER = os.Getenv("DB_DRIVER")
	DB_URL = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
}
