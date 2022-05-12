package connection

import (
	"fmt"
	"log"
	"user-product-app/config"
	"user-product-app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB	*gorm.DB
	err error
)

func Connect() {
	db_config := fmt.Sprintf("%s@tcp(%s:%s)/%s", config.GetConfig("DB_USERNAME"), config.GetConfig("DB_HOST"), config.GetConfig("DB_PORT"), config.GetConfig("DB_NAME"))
	DB, err = gorm.Open(config.GetConfig("DB_DRIVER"), db_config)

	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Server up and running")
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
}