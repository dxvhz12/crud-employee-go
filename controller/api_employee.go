package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/dxvhz12/crud-employee-go/helper"
)

func GetAllEmployeesAPI(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, nama, foto, alamat, jabatan, gaji, tanggal_masuk FROM data_karyawan")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var employees []Employee
		for rows.Next() {
			var e Employee
			rows.Scan(&e.Id, &e.Nama, &e.Foto, &e.Alamat, &e.Jabatan, &e.Gaji, &e.TanggalMasuk)
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
		row := db.QueryRow("SELECT id, nama, foto, alamat, jabatan, gaji, tanggal_masuk FROM data_karyawan WHERE id = $1", id)
		err := row.Scan(&e.Id, &e.Nama, &e.Foto, &e.Alamat, &e.Jabatan, &e.Gaji, &e.TanggalMasuk)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"employee not found"}`))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(e)
	}
} 

func UpdateEmployeeAPI(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"id wajib diisi"}`))
			return
		}

		if err := r.ParseMultipartForm(10 << 20); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"data form tidak valid"}`))
			return
		}

		nama := r.FormValue("nama")
		alamat := r.FormValue("alamat")
		jabatan := r.FormValue("jabatan")
		gaji := r.FormValue("gaji")
		tanggalMasuk := r.FormValue("tanggal_masuk")
		tanggalKeluarInput := r.FormValue("tanggal_keluar")
		currentFoto := r.FormValue("current_foto")
		status := r.FormValue("status")

		var tanggalKeluar interface{}
		if tanggalKeluarInput == "" {
			tanggalKeluar = nil
		} else {
			tanggalKeluar = tanggalKeluarInput
		}

		foto := currentFoto
		file, header, err := r.FormFile("foto")
		if err == nil {
			newFoto, err := helper.SaveUploadedFile(file, header)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error":"gagal menyimpan foto"}`))
				return
			}
			helper.DeleteUploadedFile(currentFoto)
			foto = newFoto
		} else if err != http.ErrMissingFile {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"gagal membaca file"}`))
			return
		}

		_, err = db.Exec(
			"UPDATE data_karyawan SET nama=$1, alamat=$2, jabatan=$3, gaji=$4, tanggal_masuk=$5, tanggal_keluar=$6, foto=$7, status=$8 WHERE id=$9",
			nama, alamat, jabatan, gaji, tanggalMasuk, tanggalKeluar, foto, status, id,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"gagal update data"}`))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "employee updated"})
	}
}