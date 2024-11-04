package service

import (
	"context"
	"errors"
	/* "fmt" */
	"log"
	/* "net/http" */

	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/db"
	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/model"
	"go.mongodb.org/mongo-driver/bson"
)

type Book struct{
	Title string `json:"title"`
	Author string `json:"author"`
	Stock int `json:"stock"`
	Year_released int `json:"year_released"`
	Price int `json:"price"`
}

type BookResponse struct{
	Data []*Book `json:"data"`
}

func GetAllBook() (*BookResponse, error) {
	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("Internal Server Error")
		
	}

	coll := db.MongoDB.Collection("buku")
	cur, err := coll.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("Internal Server Error")}

	var bookList []*Book

	for cur.Next(context.TODO()) {
		var book model.Book
		cur.Decode(&book)
		bookList = append(bookList, &Book{
			Title: book.Title,
			Price: book.Price,
		})

		
		
	}
	return &BookResponse{
		Data: bookList,
		}, nil
}