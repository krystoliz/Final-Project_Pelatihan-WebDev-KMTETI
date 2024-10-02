package handler

import (
	"encoding/json"
	
	"net/http"
)

type Book struct{
	Id int `json:"id"`// if its lowercase its detected as local variable
	Title string `json:"title"`
	Author string `json:"author"`
	Pages int `json:"pages"`
	Year uint `json:"year"`
}

var BookList []*Book = []*Book{
	 &Book{
		Id: 1,
        Title: "Unf/Air",
        Author: "ARTMS",
        Pages: 120,
		Year: 2024,
	},

	&Book{
		Id: 2,
        Title: "Cotton Candy",
        Author: "Loossemble",
        Pages: 120,
		Year: 2024,
	},
}

func BookHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case "GET":
            w.Header().Add("Content-Type", "application/json")
            json.NewEncoder(w).Encode(BookList)
            return

		case "POST":
			var b Book 
			json.NewDecoder(r.Body).Decode(&b)
			BookList = append(BookList, &b)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("book added successfully"))
			return
		
		case "PUT":

		case "DELETE":

        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
	}

}