package handler

import (
	"encoding/json"
	
	"net/http"
)

type Product struct{
	Id int `json:"id"`// if its lowercase its detected as local variable
	Name string `json:"name"`
	Price int `json:"price"`
	Stock uint `json:"stock"`
}

var ProdList []*Product = []*Product{
	 &Product{
		Id: 1,
        Name: "Fish",
        Price: 10000,
        Stock: 120,
	},

	&Product{
		Id: 2,
        Name: "Chips",
        Price: 15000,
        Stock: 80,
	},
}

func ProductHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		var p Product 

		json.NewDecoder(r.Body).Decode(&p)
		ProdList = append(ProdList, &p)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("product added successfully"))
		return
	}
	
	
	



	w.Header().Add("Content-Type", "application/json")
    json.NewEncoder(w).Encode(ProdList)
}