package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Status   bool   `json:"status"`
	Password string `json:"password"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error)  {
	hashed_password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashed_password)
	return
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error)  {
	current_status := user.Status
	if !current_status {
		user.Status = false
	}
	return
}
