package connection

import (
	"log"
	"user-product-app/structs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB	*gorm.DB
	err error
)

func Connect() {
	DB, err = gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/db_user_product")

	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Server up and running")
	}

	DB.AutoMigrate(&structs.Users{})
	DB.AutoMigrate(&structs.Products{})
}