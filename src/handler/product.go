package handler

import (
	"net/http"
	"encoding/json")

type Product struct{
	Id int `json:"id"`// if its lowercase its detected as local variable
	Name string `json:"name"`
	Price int `json:"price"`
	Stock uint `json:"stock"`
}

var prodList []*Product

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

	p2 := Product{
		Id: 2,
        Name: "Chips",
        Price: 15000,
        Stock: 80,
	}

	prodList = append(prodList, &p1,&p2)

	w.Header().Add("Content-Type", "application/json")
    json.NewEncoder(w).Encode(prodList)
}