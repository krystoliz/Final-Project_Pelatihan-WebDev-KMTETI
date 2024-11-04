package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	/* "fmt" */
	"log"
	/* "net/http" */

	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/db"
	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct{
	Title string `json:"title"`
	Author string `json:"author"`
	Stock int `json:"stock"`
	Year_released int `json:"year_released"`
	Price int `json:"price"`
}

type BookRequest struct{
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
	defer db.MongoDB.Client().Disconnect(context.TODO())
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
			Author: book.Author,
			Stock: book.Stock,
			Year_released: book.Year_released,
			Price: book.Price,
		})

		
		
	}
	return &BookResponse{
		Data: bookList,
		}, nil
}

func CreateBook(req io.Reader) error {
	var bookReq BookRequest
	err := json.NewDecoder(req).Decode(&bookReq)
	if err != nil {
		return errors.New("Bad Request")
	}
	db, err := db.DBConnection()
			if err != nil{
				log.Default().Println(err.Error())
				return errors.New("Internal Server Error")
			}
			defer db.MongoDB.Client().Disconnect(context.TODO())
			coll := db.MongoDB.Collection("buku")

			_, err = coll.InsertOne(context.TODO(), model.Book{
				ID: primitive.NewObjectID(),
				Title: bookReq.Title,
				Author: bookReq.Author,
				Stock: bookReq.Stock,
				Year_released: int(bookReq.Year_released),
				Price: int(bookReq.Price),
			})
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("Internal server error")
	}		
			return nil
}

