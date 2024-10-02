package handler

import (
	"encoding/json"
	
	"net/http"
)

type Product struct{
	Id int `json:"id"`// if its lowercase its detected as local variable
	Title string `json:"title"`
	Author string `json:"author"`
	Pages int `json:"pages"`
	Year uint `json:"year"`
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
	switch r.Method {
		case "GET":
            w.Header().Add("Content-Type", "application/json")
            json.NewEncoder(w).Encode(ProdList)
            return

		case "POST":
			var p Product 
			json.NewDecoder(r.Body).Decode(&p)
			ProdList = append(ProdList, &p)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("product added successfully"))
			return

		
        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
	}

}