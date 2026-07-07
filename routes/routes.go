package routes

import (
	"database/sql"
	"net/http"

	"github.com/dxvhz12/crud-employee-go/controller"
	"github.com/dxvhz12/crud-employee-go/middleware"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.HellowController())
	server.HandleFunc("/employee", controller.IndexEmployeeController(db))
	server.HandleFunc("/employee/create", controller.CreateEmployeeController(db))
	server.HandleFunc("/employee/update", controller.UpdateEmployeeController(db))
	server.HandleFunc("/employee/delete", controller.DeleteEmployeeController(db))
	server.HandleFunc("/api/employees", middleware.RequireApiKey(controller.GetAllEmployeesAPI(db)))
	server.HandleFunc("/api/employees-detail", middleware.RequireApiKey(controller.GetEmployeeByIdAPI(db)))
}