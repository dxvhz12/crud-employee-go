package routes

import (
	"database/sql"
	"net/http"

	"github.com/dxvhz12/crud-employee-go/controller"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.HellowController())
	server.HandleFunc("/employee", controller.IndexEmployeeController(db))
	server.HandleFunc("/employee/create", controller.CreateEmployeeController(db))
	server.HandleFunc("/employee/update", controller.UpdateEmployeeController(db))
	server.HandleFunc("/employee/delete", controller.DeleteEmployeeController(db))
}