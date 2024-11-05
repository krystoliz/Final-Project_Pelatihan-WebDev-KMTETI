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
	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct{
	Title string `json:"title"`
	Author string `json:"author"`
	/* Stock int `json:"stock"`
	Year_released int `json:"year_released"` */
	Price int `json:"price"`
}

type BookRequest struct{
	Title string `json:"title"`
	Author string `json:"author"`
	Stock int `json:"stock"`
	Year_released int `json:"year_released"`
	Price int `json:"price"`
}

type UpdateBookRequest struct{
	Title string `json:"title"`
	Price int `json:"price"`
	Stock int `json:"stock"`
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
		var book model.ShowAllBook
		cur.Decode(&book)
		bookList = append(bookList, &Book{
			Title: book.Title,
			Author: book.Author,
			
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

func GetBookByTitle(title string) (*BookRequest, error) {
    // Establish database connection
    db, err := db.DBConnection()
    if err != nil {
        log.Default().Println(err.Error())
        return nil, errors.New("Internal Server Error")
    }
    defer db.MongoDB.Client().Disconnect(context.TODO())

    // Get collection reference
    coll := db.MongoDB.Collection("buku")

    // Create filter for the query
    filter := bson.M{"title": title}

    // Create a variable to store the result
    var book model.Book
    err = coll.FindOne(context.TODO(), filter).Decode(&book)
    
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, errors.New("Book not found")
        }
        log.Default().Println(err.Error())
        return nil, errors.New("Internal Server Error")
    }

    // Convert to response type
    response := &BookRequest{
       
        Title:         book.Title,
        Author:        book.Author,
        Stock:         book.Stock,
        Year_released: book.Year_released,
        Price:         book.Price,
    }

    return response, nil
}

func UpdateBook(req io.Reader) error{
	var updateReq UpdateBookRequest
	err := json.NewDecoder(req).Decode(&updateReq)
	if err != nil {
		return errors.New("Bad Request")
	}
	//validate request
	if updateReq.Title == ""{
		return errors.New("Title is required")
	}
	db, err := db.DBConnection()
    if err != nil {
        log.Default().Println(err.Error())
        return errors.New("Internal Server Error")
    }
    defer db.MongoDB.Client().Disconnect(context.TODO())

    coll := db.MongoDB.Collection("buku")

    // Create filter and update document
    filter := bson.M{"title": updateReq.Title}
    update := bson.M{
        "$set": bson.M{
            "price": updateReq.Price,
            "stock": updateReq.Stock,
        },
    }

    result, err := coll.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        log.Default().Println(err.Error())
        return errors.New("Internal Server Error")
    }

    if result.MatchedCount == 0 {
        return errors.New("Book not found")
    }

    return nil
}

func DeleteBook(title string) error {
    if title == "" {
        return errors.New("Title is required")
    }

    db, err := db.DBConnection()
    if err != nil {
        log.Default().Println(err.Error())
        return errors.New("Internal Server Error")
    }
    defer db.MongoDB.Client().Disconnect(context.TODO())

    coll := db.MongoDB.Collection("buku")

    // Create filter for the specific book
    filter := bson.M{"title": title}

    // Delete the document
    result, err := coll.DeleteOne(context.TODO(), filter)
    if err != nil {
        log.Default().Println(err.Error())
        return errors.New("Internal Server Error")
    }

    if result.DeletedCount == 0 {
        return errors.New("Book not found")
    }

    return nil
}