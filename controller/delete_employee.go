package controller

import (
	"database/sql"
	"net/http"

	"github.com/dxvhz12/crud-employee-go/helper"
)

func DeleteEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		var foto sql.NullString
		db.QueryRow("SELECT foto FROM data_karyawan WHERE id = $1", id).Scan(&foto)

		_, err := db.Exec("DELETE FROM data_karyawan WHERE id = $1", id)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		helper.DeleteUploadedFile(foto.String)

		http.Redirect(w, r, "/employee", http.StatusMovedPermanently)
	}
}
