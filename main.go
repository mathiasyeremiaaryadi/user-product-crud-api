package main

import (
	"user-product-app/connection"
	"user-product-app/routes"
)

func main() {
	connection.Connect()
	routes.APIRoute()
}