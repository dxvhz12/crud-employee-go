package controller

import (
	"html/template"

	"github.com/dxvhz12/crud-employee-go/helper"
)

var funcMap = template.FuncMap{
	"rupiah":    helper.FormatRupiah,
	"tanggalID": helper.FormatTanggal,
	"initial":   helper.Initial,
	"dateISO":   helper.FormatDateISO,
}