package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/dxvhz12/crud-employee-go/helper"
)

func GetAllEmployeesAPI(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, nama, foto, alamat, jabatan, gaji, tanggal_masuk, tanggal_keluar, status FROM data_karyawan")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var employees []Employee
		for rows.Next() {
			var e Employee
			var tanggalKeluar sql.NullString
			if err := rows.Scan(&e.Id, &e.Nama, &e.Foto, &e.Alamat, &e.Jabatan, &e.Gaji, &e.TanggalMasuk, &tanggalKeluar, &e.Status); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			e.TanggalKeluar = tanggalKeluar.String
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
		var tanggalKeluar sql.NullString
		row := db.QueryRow("SELECT id, nama, foto, alamat, jabatan, gaji, tanggal_masuk, tanggal_keluar, status FROM data_karyawan WHERE id = $1", id)
		err := row.Scan(&e.Id, &e.Nama, &e.Foto, &e.Alamat, &e.Jabatan, &e.Gaji, &e.TanggalMasuk, &tanggalKeluar, &e.Status)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"employee not found"}`))
			return
		}
		e.TanggalKeluar = tanggalKeluar.String

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(e)
	}
}

func CreateEmployeeAPI(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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

		var foto string
		file, header, err := r.FormFile("foto")
		if err == nil {
			foto, err = helper.SaveUploadedFile(file, header)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error":"gagal menyimpan foto"}`))
				return
			}
		} else if err != http.ErrMissingFile {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"gagal membaca file"}`))
			return
		}

		var id string
		err = db.QueryRow(
			"INSERT INTO data_karyawan (nama, alamat, jabatan, gaji, tanggal_masuk, foto, status) VALUES ($1, $2, $3, $4, $5, $6, 'aktif') RETURNING id",
			nama, alamat, jabatan, gaji, tanggalMasuk, foto,
		).Scan(&id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"gagal menyimpan data"}`))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "employee created", "id": id})
	}
}

func DeleteEmployeeAPI(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"id wajib diisi"}`))
			return
		}

		var foto sql.NullString
		db.QueryRow("SELECT foto FROM data_karyawan WHERE id = $1", id).Scan(&foto)

		_, err := db.Exec("DELETE FROM data_karyawan WHERE id = $1", id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"gagal menghapus data"}`))
			return
		}

		helper.DeleteUploadedFile(foto.String)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "employee deleted"})
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