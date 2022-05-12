package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"user-product-app/connection"
	"user-product-app/models"
	"user-product-app/services"

	"golang.org/x/crypto/bcrypt"
)


func Login(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var user_model models.User
	var res services.JSONService

	json.Unmarshal(payloads, &user_model)
	input_pass := user_model.Password
	connection.DB.Where("email=?", &user_model.Email).Find(&user_model)
	err := bcrypt.CompareHashAndPassword([]byte(user_model.Password), []byte(input_pass))

	if err != nil || user_model.ID == 0 {
		res = services.JSONService{Code: 200, Data: nil, Message: "Invalid email or password"}
	} else {
		res = services.JSONService{Code: 200, Data: user_model, Message: "Authenticated"}
	}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

