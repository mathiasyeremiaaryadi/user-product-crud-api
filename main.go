package main

import (
	"user-product-app/connection"
	"user-product-app/handlers"
)

func main() {
	connection.Connect()
	handlers.HandleReq()
}