package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Employee struct {
	Id           string  `json:"id"`
	Nama         string  `json:"nama"`
	Foto         string  `json:"foto"`
	Alamat       string  `json:"alamat"`
	Jabatan      string  `json:"jabatan"`
	Gaji         float64 `json:"gaji"`
	TanggalMasuk string  `json:"tanggal_masuk"`
	TanggalKeluar string `json:"tanggal_keluar"`
	Status       string  `json:"status"`
}

func IndexEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, nama, alamat, jabatan, gaji, tanggal_masuk, tanggal_keluar, foto, status FROM data_karyawan")
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var employees []Employee
		for rows.Next() {
			var employee Employee
			var tanggalKeluar, foto sql.NullString

			err = rows.Scan(
				&employee.Id,
				&employee.Nama,
				&employee.Alamat,
				&employee.Jabatan,
				&employee.Gaji,
				&employee.TanggalMasuk,
				&tanggalKeluar,
				&foto,
				&employee.Status,
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			employee.TanggalKeluar = tanggalKeluar.String
			employee.Foto = foto.String
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
