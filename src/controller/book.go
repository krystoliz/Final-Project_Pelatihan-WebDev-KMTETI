package controller

import (
	/* "context" */
	"encoding/json"
	/* "fmt" */
	"net/http"
	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/service"
	/* "github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/db"
	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" */
)

type Book struct{
	// if its lowercase its detected as local variable
	Title string `json:"title"`
	Author string `json:"author"`
	Stock int `json:"stock"`
	Year_released uint `json:"year_released"`
	Price uint `json:"price"`
}

var BookList []*Book = []*Book{
	 &Book{
		
        Title: "Unf/Air",
        Author: "ARTMS",
        Stock: 120,
		Year_released: 2024,
		Price: 100000,
	},

	&Book{
		
        Title: "Cotton Candy",
        Author: "Loossemble",
        Stock: 120,
		Year_released: 2024,
		Price: 100000,
	},
}

func BookController(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		// check if theres a title query parameter
		title := r.URL.Query().Get("title")

		if title != "" {
			//Get specific book by name
			book, err := service.GetBookByTitle(title)
			if err != nil {
				if err.Error()== "Book bot found"{
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}



		//Get all book if no title parameter
		data, err := service.GetAllBook()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
		/* w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bookList) */
		return



	case "POST":
		err := service.CreateBook(r.Body)
		if err != nil {
			if err.Error() == "bad request"{
				http.Error(w, err.Error(), http.StatusBadRequest)

			}
			http.Error(w, "internal server error" , http.StatusInternalServerError)
			
		}
		defer r.Body.Close()
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Book has been created successfully")
	case "PUT":

	case "DELETE":

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
}
	w.Write([]byte("Hello from book"))
}



func NewProductHandler(w http.ResponseWriter, r *http.Request){
	
  w.Write([]byte("connection success"))
}	