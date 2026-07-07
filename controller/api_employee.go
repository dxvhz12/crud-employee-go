package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetAllEmployeesAPI(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, nama, alamat, jabatan, gaji, tanggal_masuk FROM data_karyawan")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var employees []Employee
		for rows.Next() {
			var e Employee
			rows.Scan(&e.Id, &e.Nama, &e.Alamat, &e.Jabatan, &e.Gaji, &e.TanggalMasuk)
			employees = append(employees, e)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employees)
	}
}

func GetEmployeeByIdAPI(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		var e Employee
		row := db.QueryRow("SELECT id, nama, alamat, jabatan, gaji, tanggal_masuk FROM data_karyawan WHERE id = ?", id)
		err := row.Scan(&e.Id, &e.Nama, &e.Alamat, &e.Jabatan, &e.Gaji, &e.TanggalMasuk)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"employee not found"}`))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(e)
	}
}