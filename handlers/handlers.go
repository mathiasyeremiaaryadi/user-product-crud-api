package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleReq() {
	log.Println("Start development server localhost:8000")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users", GetUsers).Queries(
		"limit", "{limit}",
		"offset", "{offset}",
	).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	router.HandleFunc("/products", GetProducts).Queries(
		"limit", "{limit}",
		"offset", "{offset}",
	).Methods("GET")
	router.HandleFunc("/products", CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/products/{id}", UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}