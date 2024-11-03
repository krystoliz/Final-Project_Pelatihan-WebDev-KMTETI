package handler

import "net/http"

func BookHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from book"))
}