package helper

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

const uploadDir = "public/uploads"

// SaveUploadedFile menyimpan file upload ke folder uploads,
// mengembalikan NAMA FILE saja (bukan full path) untuk disimpan di DB.
func SaveUploadedFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	defer file.Close()

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", err
	}

	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	dst := filepath.Join(uploadDir, filename)

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		return "", err
	}

	return filename, nil
}

// DeleteUploadedFile menghapus file lama dari disk (dipanggil saat foto diganti/employee dihapus)
func DeleteUploadedFile(filename string) {
	if filename == "" {
		return
	}
	os.Remove(filepath.Join(uploadDir, filename))
}