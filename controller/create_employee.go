package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/dxvhz12/crud-employee-go/helper"
)

func CreateEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := r.ParseMultipartForm(10 << 20) // maksimal 10MB
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			nama := r.FormValue("nama")
			alamatku := r.FormValue("address")
			jabatan := r.FormValue("jabatan")
			gaji := r.FormValue("gaji")
			tanggalMasuk := r.FormValue("tanggal_masuk")
			status := "aktif"

			var foto string
			file, header, err := r.FormFile("foto")
			if err == nil {
				foto, err = helper.SaveUploadedFile(file, header)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			} else if err != http.ErrMissingFile {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			_, err = db.Exec(
				"INSERT INTO data_karyawan (nama, alamat, jabatan, gaji, tanggal_masuk, foto, status) VALUES ($1, $2, $3, $4, $5, $6, $7)",
				nama, alamatku, jabatan, gaji, tanggalMasuk, foto, status,
			)
			
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
