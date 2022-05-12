package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"user-product-app/connection"
	"user-product-app/models"
	"user-product-app/services"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var limit int
	var offset int
	
	keys, ok := r.URL.Query()["limit"]
	if ok && len(keys[0]) >= 1 {
		limit, _ = strconv.Atoi(keys[0])
	}

	keys, ok = r.URL.Query()["offset"]
	if ok && len(keys[0]) >= 1 {
		offset, _ = strconv.Atoi(keys[0])
	}

	if limit == 0 {
		limit = 10
	}

	user_model := []models.User{}

	connection.DB.Limit(limit).Offset(offset).Find(&user_model)

	res := services.JSONService{Code: 200, Data: user_model, Message: "User has successfully retrieved"}
	resuts, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resuts)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]

	user_model := models.User{}
	connection.DB.First(&user_model, user_id)

	res := services.JSONService{Code: 200, Data: user_model, Message: "User has found"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var user_model models.User

	json.Unmarshal(payloads, &user_model)
	connection.DB.Create(&user_model)

	res := services.JSONService{Code: 200, Data: user_model, Message: "User has successfully created"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]
	payloads, _ := ioutil.ReadAll(r.Body)

	var user_model models.User
	connection.DB.First(&user_model, user_id)

	json.Unmarshal(payloads, &user_model)
	connection.DB.Save(&user_model)

	res := services.JSONService{Code: 200, Data: user_model, Message: "User has successfully updated"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]

	var user_model models.User
	connection.DB.First(&user_model, user_id)
	connection.DB.Delete(&user_model)

	res := services.JSONService{Code: 200, Data: user_model, Message: "User has successfully deleted"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}