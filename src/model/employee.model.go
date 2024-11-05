package model

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkStatus string

const (
    Contract WorkStatus = "Contract"
    FullTime WorkStatus = "Full-time"
)

type Employee struct {
    ID              primitive.ObjectID `bson:"_id"`
    Name            string            `bson:"name"`
    NIK             string            `bson:"nik"`
    Education_level string            `bson:"education_level"`
    Date_started    time.Time         `bson:"date_started"`
    Work_status     WorkStatus        `bson:"work_status"`
}
