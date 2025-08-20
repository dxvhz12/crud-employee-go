package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func CreateEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			nama := r.Form["nama"][0]
			alamatku := r.Form["address"][0]
			jabatan := r.Form["jabatan"][0]
			gaji := r.Form["gaji"][0]
			tanggal_masuk := r.Form["tanggal_masuk"][0]
			
			_, err := db.Exec("INSERT INTO data_karyawan (nama, alamat, jabatan, gaji, tanggal_masuk) VALUES (?, ?, ?, ?, ?)", nama, alamatku, jabatan, gaji, tanggal_masuk)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			
			http.Redirect(w, r, "/employee", http.StatusMovedPermanently)
			return 
		} else if r.Method == "GET" {
			fp := filepath.Join("views", "create.html")

			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, nil)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
