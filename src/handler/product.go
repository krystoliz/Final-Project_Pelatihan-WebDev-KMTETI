package handler

import "net/http"

type Product struct{
	Id int// if its lowercase its detected as local variable
	Name string
	Price int
	Stock uint
}

func ProductHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed) 
		w.Write([]byte("Method Not Allowed"))
		return
	}
	
    w.Write([]byte("Hello from /api/product"))
}