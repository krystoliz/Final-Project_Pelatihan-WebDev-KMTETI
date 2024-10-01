package handler

import "net/http"

type Product struct{
	Id int `json:"id"`// if its lowercase its detected as local variable
	Name string `json:"name"`
	Price int `json:"price"`
	Stock uint `json:"stock"`
}

func ProductHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed) 
		w.Write([]byte("Method Not Allowed"))
		return
	}
	
	p1 := Product{
		Id: 1,
        Name: "Fish",
        Price: 10000,
        Stock: 120,
	}
    w.Write([]byte("Hello from /api/product"))
}