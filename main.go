package main

import (
	"log"
	"net/http"
	"project-crud/config"
	"project-crud/controllers/categorycontroller"
	"project-crud/controllers/homecontroller"
	"project-crud/controllers/productcontroller"
)

func main() {
	config.ConnectDB()

	// 1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)

	// 3. Products
	http.HandleFunc("/products", productcontroller.Index)
	
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}