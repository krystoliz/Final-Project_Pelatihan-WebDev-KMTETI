package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/db"
	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
			
			var bookList []*model.Book

			for cur.Next(context.TODO()){
				var prod model.Book
				err:= cur.Decode(&prod)
				if err!= nil{
					http.Error(w, "Error decoding data", http.StatusInternalServerError)
					return
				}
				
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

			db, err := db.DBConnection()
			if err != nil{
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			coll := db.MongoDB.Collection("buku")

			_, err = coll.InsertOne(context.TODO(), model.Book{
				ID: primitive.NewObjectID(),
				Title: b.Title,
				Author: b.Author,
				Stock: b.Stock,
				Year_released: int(b.Year_released),
				Price: int(b.Price),
			})
			
			if err != nil{
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("book added successfully"))
		
		case "PUT":

		case "DELETE":

        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
			return
	}

}



func NewProductHandler(w http.ResponseWriter, r *http.Request){
	
  w.Write([]byte("connection success"))
}	