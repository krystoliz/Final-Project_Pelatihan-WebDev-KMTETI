package handler

import (
	"context"
	"encoding/json"
	/* "fmt" */

	"net/http"

	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/db"
	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/model"
	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/service"
	/* "go.mongodb.org/mongo-driver/bson" */
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BookHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
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
		var b model.Book 
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
	w.Write([]byte("Hello from book"))
}