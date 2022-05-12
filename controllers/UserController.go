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

	db_user_product := []models.User{}

	connection.DB.Limit(limit).Offset(offset).Find(&db_user_product)

	res := services.JSONService{Code: 200, Data: db_user_product, Message: "User has successfully retrieved"}
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

	db_user_product := models.User{}
	connection.DB.First(&db_user_product, user_id)

	res := services.JSONService{Code: 200, Data: db_user_product, Message: "User has found"}

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

	var db_user_product models.User
	json.Unmarshal(payloads, &db_user_product)
	connection.DB.Create(&db_user_product)

	res := services.JSONService{Code: 200, Data: db_user_product, Message: "User has successfully created"}

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

	var db_user_product models.User
	connection.DB.First(&db_user_product, user_id)

	json.Unmarshal(payloads, &db_user_product)
	connection.DB.Model(&db_user_product).Update(db_user_product)
	if !db_user_product.Status {
		connection.DB.Model(&db_user_product).Update(map[string]interface{}{"status": false})
	}

	res := services.JSONService{Code: 200, Data: db_user_product, Message: "User has successfully updated"}

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

	var db_user_product models.User
	connection.DB.First(&db_user_product, user_id)
	connection.DB.Delete(&db_user_product)

	res := services.JSONService{Code: 200, Data: db_user_product, Message: "User has successfully deleted"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}