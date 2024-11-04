package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID            primitive.ObjectID `bson:"_id,emitonempty"`
	Title         string             `bson:"title"`
	Author        string             `bson:"author"`
	Stock         int                `bson:"stock"`
	Year_released int                `bson:"year_released"`
	Price         int                `bson:"price"`
}
