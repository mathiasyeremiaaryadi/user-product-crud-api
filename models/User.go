package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username" validate:"required,alpha"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"numeric"`
	Role     string `json:"role" validate:"required,alpha"`
	Status   bool   `json:"status" validate:"required,boolean"`
	Password string `json:"password" validate:"required"`
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
