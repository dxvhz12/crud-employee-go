package main

import (
	"net/http"

	"github.com/dxvhz12/crud-employee-go/database"
	"github.com/dxvhz12/crud-employee-go/routes"
)

func main() {
	db := database.InitDatabase()
	
	server := http.NewServeMux()
	routes.MapRoutes(server, db)

	http.ListenAndServe(":9090", server)
}