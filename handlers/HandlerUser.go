package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"user-product-app/connection"
	"user-product-app/structs"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var limit interface{}
	var offset interface{}

	limit = vars["limit"]
	offset = vars["offset"]

	if limit == ""  {
		limit = 10
	}

	if offset == "" {
		offset = 0
	}

	db_user_product := []structs.Users{}

	connection.DB.Limit(limit).Offset(offset).Find(&db_user_product)

	res := structs.Result{Code: 200, Data: db_user_product, Message: "User has successfully retrieve"}
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

	db_user_product := structs.Users{}
	connection.DB.First(&db_user_product, user_id)

	res := structs.Result{Code: 200, Data: db_user_product, Message: "User has found"}

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

	var db_user_product structs.Users
	json.Unmarshal(payloads, &db_user_product)
	connection.DB.Create(&db_user_product)

	res := structs.Result{Code: 200, Data: db_user_product, Message: "User has successfully created"}

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

	var db_user_product structs.Users
	connection.DB.First(&db_user_product, user_id)

	json.Unmarshal(payloads, &db_user_product)
	connection.DB.Model(&db_user_product).Update(db_user_product)

	res := structs.Result{Code: 200, Data: db_user_product, Message: "User has successfully updated"}

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

	var db_user_product structs.Users
	connection.DB.First(&db_user_product, user_id)
	connection.DB.Delete(&db_user_product)

	res := structs.Result{Code: 200, Data: db_user_product, Message: "User has successfully deleted"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}