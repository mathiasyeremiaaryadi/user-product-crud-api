package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func (product *Product) BeforeUpdate(tx *gorm.DB) (err error)  {
	current_status := product.Status
	if !current_status {
		product.Status = false
	}
	return
}