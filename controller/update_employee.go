package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/dxvhz12/crud-employee-go/helper"
)

func UpdateEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			id := r.URL.Query().Get("id")
			err := r.ParseMultipartForm(10 << 20)
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
			tanggalKeluarInput := r.FormValue("tanggal_keluar")
			currentFoto := r.FormValue("current_foto")

			// status ditentukan otomatis dari ada/tidaknya tanggal_keluar,
			// bukan dari input user
			var tanggalKeluar interface{}
			var status string
			if tanggalKeluarInput == "" {
				tanggalKeluar = nil
				status = "aktif"
			} else {
				tanggalKeluar = tanggalKeluarInput
				status = "nonaktif"
			}

			foto := currentFoto
			file, header, err := r.FormFile("foto")
			if err == nil {
				newFoto, err := helper.SaveUploadedFile(file, header)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				helper.DeleteUploadedFile(currentFoto)
				foto = newFoto
			} else if err != http.ErrMissingFile {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			_, err = db.Exec(
				"UPDATE data_karyawan SET nama=$1, alamat=$2, jabatan=$3, gaji=$4, tanggal_masuk=$5, tanggal_keluar=$6, foto=$7, status=$8 WHERE id=$9",
				nama, alamatku, jabatan, gaji, tanggalMasuk, tanggalKeluar, foto, status, id,
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/employee", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			id := r.URL.Query().Get("id")

			var employee Employee
			var tanggalKeluar, foto sql.NullString

			row := db.QueryRow("SELECT nama, alamat, jabatan, gaji, tanggal_masuk, tanggal_keluar, foto, status FROM data_karyawan WHERE id = $1", id)
			err := row.Scan(
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
			employee.Id = id
			employee.TanggalKeluar = tanggalKeluar.String
			employee.Foto = foto.String

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