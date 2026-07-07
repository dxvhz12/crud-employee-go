package main

import (
	"log"
	"net/http"

	"github.com/dxvhz12/crud-employee-go/database"
	"github.com/dxvhz12/crud-employee-go/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file tidak ditemukan, pakai environment variable sistem")
	}

	db := database.InitDatabase()
	
	server := http.NewServeMux()
	routes.MapRoutes(server, db)

	http.ListenAndServe(":9090", server)
}