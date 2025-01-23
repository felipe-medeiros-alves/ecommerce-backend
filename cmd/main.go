package main

import (
	"ecommerce-backend/config"
	"ecommerce-backend/routes"
	"log"
	"net/http"
)

func main() {
	config.InitDB()
	config.InitRedis()

	router := routes.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
