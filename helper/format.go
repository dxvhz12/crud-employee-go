package helper

import (
	"strconv"
	"strings"
	"time"
)

// FormatRupiah: 5000000 -> "5.000.000"
func FormatRupiah(n float64) string {
	in := strconv.FormatFloat(n, 'f', 0, 64)
	neg := strings.HasPrefix(in, "-")
	if neg {
		in = in[1:]
	}

	var out []byte
	l := len(in)
	for i := 0; i < l; i++ {
		if i != 0 && (l-i)%3 == 0 {
			out = append(out, '.')
		}
		out = append(out, in[i])
	}

	result := string(out)
	if neg {
		result = "-" + result
	}
	return result
}

// FormatTanggal: "2025-12-24" atau "2025-12-24 00:00:00" -> "24-12-2025"
func FormatTanggal(tanggal string) string {
	layouts := []string{
		"2006-01-02T15:04:05Z",
		"2006-01-02 15:04:05",
		"2006-01-02",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, tanggal); err == nil {
			return t.Format("02-01-2006")
		}
	}

	return tanggal // fallback kalau semua format gagal di-parse
}