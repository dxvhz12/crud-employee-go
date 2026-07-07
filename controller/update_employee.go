package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func UpdateEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			id := r.URL.Query().Get("id")
			r.ParseForm()

			nama := r.Form["nama"][0]
			alamatku := r.Form["address"][0]
			jabatan := r.Form["jabatan"][0]
			gaji := r.Form["gaji"][0]
			tanggal_masuk := r.Form["tanggal_masuk"][0]

			_, err := db.Exec("UPDATE data_karyawan SET nama=?, alamat=?, jabatan=?, gaji=?, tanggal_masuk=? WHERE id=?", nama, alamatku, jabatan, gaji, tanggal_masuk, id)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/employee", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			id := r.URL.Query().Get("id")
			row := db.QueryRow("SELECT nama, alamat, jabatan, gaji, tanggal_masuk FROM data_karyawan WHERE id = ?", id)
			if row.Err() != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var employee Employee
			err := row.Scan(
				&employee.Nama,
				&employee.Alamat,
				&employee.Jabatan,
				&employee.Gaji,
				&employee.TanggalMasuk,
			)
			employee.Id = id
			if err != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			fp := filepath.Join("views", "update.html")

			tmpl, err := template.New("update.html").Funcs(funcMap).ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			data := make(map[string]any)
			data["employee"] = employee

			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
