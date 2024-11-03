package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/db"
	"go.mongodb.org/mongo-driver/bson"
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

func BookHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
		case "GET":
			db, err := db.DBConnection()
			if err != nil{
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			coll := db.MongoDB.Collection("buku")
			cur, err := coll.Find(context.TODO(), bson.D{})
			if err != nil{
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			var prod BookModel
			var bookList []*BookModel

			for cur.Next(context.TODO()){
				cur.Decode(&prod)
				fmt.Println(prod)
				bookList = append(bookList, &prod)
			}

            w.Header().Add("Content-Type", "application/json")
            json.NewEncoder(w).Encode(bookList)
            return



		case "POST":
			var b Book 
			err := json.NewDecoder(r.Body).Decode(&b)
			if err != nil{
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
		
		case "PUT":

		case "DELETE":

        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
	}

}

type BookModel struct{
	ID string `bson:"_id,emitonempty"`
	Title string `bson:"title"`
	Author string `bson:"author"`
	Stock int `bson:"stock"`
	Year_released int `bson:"year_released"`
	Price int `bson:"price"`
}

func NewProductHandler(w http.ResponseWriter, r *http.Request){
	
  w.Write([]byte("connection success"))
}	