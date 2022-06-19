package api

import (
	"api/routes"
	"config"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Run() {
	config.Load()
	ConnectDB()
	fmt.Printf("Listening PORT:%d", config.PORT)
	Listen(config.PORT)
}

func Listen(port int) {
	r := routes.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func ConnectDB() {
	_, err := gorm.Open(config.DB_DRIVER, config.DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	fmt.Println("Successfully Connected to MySQL database")
}
