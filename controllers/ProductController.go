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

func GetProducts(w http.ResponseWriter, r *http.Request) {
	product_model := []models.Product{}
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

	connection.DB.Limit(limit).Offset(offset).Find(&product_model)
	res := services.JSONService{Code: 200, Data: product_model, Message: "Products has successfully retrieved"}
	resuts, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resuts)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	product_id := vars["id"]
	product_model := models.Product{}
	var res services.JSONService

	if err := connection.DB.First(&product_model, product_id); err.Error != nil {
		res = services.JSONService{Code: 404, Data: nil, Message: "Product not found"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
	} else {
		res = services.JSONService{Code: 200, Data: product_model, Message: "Product has found"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

	result, _ := json.Marshal(res)
	w.Write(result)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)
	var product_model models.Product

	json.Unmarshal(payloads, &product_model)
	connection.DB.Create(&product_model)

	res := services.JSONService{Code: 201, Data: nil, Message: "Product has successfully created"}
	result, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	product_id := vars["id"]
	payloads, _ := ioutil.ReadAll(r.Body)
	var res services.JSONService
	var product_model models.Product

	if err:= connection.DB.First(&product_model, product_id); err.Error != nil {
		res = services.JSONService{Code: 404, Data: nil, Message: "Product not found"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
	} else {
		json.Unmarshal(payloads, &product_model)
		connection.DB.Save(&product_model)
		res = services.JSONService{Code: 201, Data: nil, Message: "Product has successfully updated"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}

	result, _ := json.Marshal(res)
	w.Write(result)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	product_id := vars["id"]
	var res services.JSONService
	var product_model models.Product

	if err := connection.DB.First(&product_model, product_id); err.Error != nil {
		res = services.JSONService{Code: 404, Data: nil, Message: "Product not found"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
	} else {
		connection.DB.Delete(&product_model)
		res = services.JSONService{Code: 200, Data: nil, Message: "Product has successfully deleted"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}

	result, _ := json.Marshal(res)
	w.Write(result)
}