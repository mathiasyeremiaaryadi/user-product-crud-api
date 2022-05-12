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
	user_model := []models.User{}
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

	connection.DB.Limit(limit).Offset(offset).Find(&user_model)
	res := services.JSONService{Code: 200, Data: user_model, Message: "User has successfully retrieved"}
	resuts, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resuts)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]
	user_model := models.User{}
	var res services.JSONService

	if err := connection.DB.First(&user_model, user_id); err.Error != nil {
		res = services.JSONService{Code: 404, Data: nil, Message: "User not found"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
	} else {
		res = services.JSONService{Code: 200, Data: user_model, Message: "User has found"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

	result, _ := json.Marshal(res)
	w.Write(result)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)
	var user_model models.User

	json.Unmarshal(payloads, &user_model)
	connection.DB.Create(&user_model)

	res := services.JSONService{Code: 201, Data: nil, Message: "User has successfully created"}
	result, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]
	payloads, _ := ioutil.ReadAll(r.Body)
	var res services.JSONService

	var user_model models.User
	if err := connection.DB.First(&user_model, user_id); err.Error != nil {
		res = services.JSONService{Code: 404, Data: nil, Message: "User not found"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.Unmarshal(payloads, &user_model)
		connection.DB.Save(&user_model)
		res = services.JSONService{Code: 201, Data: nil, Message: "User has successfully updated"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}

	result, _ := json.Marshal(res)
	w.Write(result)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]
	var user_model models.User
	var res services.JSONService

	if err := connection.DB.First(&user_model, user_id); err.Error != nil {
		res = services.JSONService{Code: 404, Data: nil, Message: "User not found"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	} else {
		connection.DB.Delete(&user_model)
		res = services.JSONService{Code: 200, Data: nil, Message: "User has successfully deleted"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
	
	result, _ := json.Marshal(res)
	w.Write(result)
}