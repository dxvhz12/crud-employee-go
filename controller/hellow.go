package controller

import "net/http"

func HellowController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hellow broo"))
	}
}