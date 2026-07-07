package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Employee struct {
	Id           string
	Nama         string
	Alamat       string
	Jabatan      string
	Gaji         float64
	TanggalMasuk string
}

func IndexEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, nama, alamat, jabatan, gaji, tanggal_masuk FROM data_karyawan")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var employees []Employee
		for rows.Next() {
			var employee Employee

			err = rows.Scan(
				&employee.Id,
				&employee.Nama,
				&employee.Alamat,
				&employee.Jabatan,
				&employee.Gaji,
				&employee.TanggalMasuk,
			)

			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			
			employees = append(employees, employee)
		}

		fp := filepath.Join("views", "index.html")

		tmpl, err := template.New("index.html").Funcs(funcMap).ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["employees"] = employees

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
