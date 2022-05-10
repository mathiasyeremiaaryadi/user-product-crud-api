package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"user-product-app/connection"
	"user-product-app/structs"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
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

	db_user_product := []structs.Products{}

	connection.DB.Limit(limit).Offset(offset).Find(&db_user_product)

	res := structs.Result{Code: 200, Data: db_user_product, Message: "Products has successfully retrieve"}
	resuts, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resuts)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	product_id := vars["id"]

	db_user_product := structs.Products{}
	connection.DB.First(&db_user_product, product_id)

	res := structs.Result{Code: 200, Data: db_user_product, Message: "Product has found"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var db_user_product structs.Products
	json.Unmarshal(payloads, &db_user_product)
	connection.DB.Create(&db_user_product)

	res := structs.Result{Code: 200, Data: db_user_product, Message: "Product has successfully created"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	product_id := vars["id"]
	payloads, _ := ioutil.ReadAll(r.Body)

	var db_user_product structs.Products
	connection.DB.First(&db_user_product, product_id)

	json.Unmarshal(payloads, &db_user_product)
	connection.DB.Model(&db_user_product).Updates(db_user_product)

	var data map[string]interface{}
	res := structs.Result{Code: 200, Data: data, Message: "User has successfully updated"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	product_id := vars["id"]

	var db_user_product structs.Products
	connection.DB.First(&db_user_product, product_id)
	connection.DB.Delete(&db_user_product)

	var data map[string]interface{}
	res := structs.Result{Code: 200, Data: data, Message: "Product has successfully updated"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}