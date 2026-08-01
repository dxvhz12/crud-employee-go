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
	server.HandleFunc("GET /api/employees", middleware.VerifyJWT(controller.GetAllEmployeesAPI(db)))
	server.HandleFunc("GET /api/employees-detail", middleware.VerifyJWT(controller.GetEmployeeByIdAPI(db)))
	server.HandleFunc("POST /api/employee", middleware.VerifyJWT(controller.CreateEmployeeAPI(db)))
	server.HandleFunc("PUT /api/employee", middleware.VerifyJWT(controller.UpdateEmployeeAPI(db)))
	server.HandleFunc("DELETE /api/employee", middleware.VerifyJWT(controller.DeleteEmployeeAPI(db)))

	// serve file foto yang diupload
	server.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("public/uploads"))))
}