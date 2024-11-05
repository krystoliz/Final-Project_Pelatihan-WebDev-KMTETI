package service

import (
    "context"
    "encoding/json"
    "errors"
    "io"
    "log"
    "time"
    "github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/model"
    "github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/db"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeRequest struct {
    Name            string         `json:"name"`
    NIK             string         `json:"nik"`
    EducationLevel  string         `json:"educationLevel"`
    DateStarted     string         `json:"dateStarted"`  // Accepting date as string in format "2006-01-02"
    WorkStatus      model.WorkStatus `json:"workStatus"`
}

func CreateEmployee(req io.Reader) error {
    var empReq EmployeeRequest
    err := json.NewDecoder(req).Decode(&empReq)
    if err != nil {
        return errors.New("Bad Request")
    }

    // Validate required fields
    if empReq.Name == "" || empReq.NIK == "" {
        return errors.New("Name and NIK are required")
    }

    // Validate work status
    if empReq.WorkStatus != model.Contract && empReq.WorkStatus != model.FullTime {
        return errors.New("Invalid work status. Must be either 'Contract' or 'Full-time'")
    }

    // Parse date string to time.Time
    dateStarted, err := time.Parse("2006-01-02", empReq.DateStarted)
    if err != nil {
        return errors.New("Invalid date format. Use YYYY-MM-DD")
    }

    // Connect to database
    db, err := db.DBConnection()
    if err != nil {
        log.Default().Println(err.Error())
        return errors.New("Internal Server Error")
    }
    defer db.MongoDB.Client().Disconnect(context.TODO())

    // Get collection
    coll := db.MongoDB.Collection("employee")

    // Create new employee document
    employee := model.Employee{
        ID:              primitive.NewObjectID(),
        Name:            empReq.Name,
        NIK:             empReq.NIK,
        Education_level: empReq.EducationLevel,
        Date_started:    dateStarted,
        Work_status:     empReq.WorkStatus,
    }

    // Insert into database
    _, err = coll.InsertOne(context.TODO(), employee)
    if err != nil {
        log.Default().Println(err.Error())
        return errors.New("Internal Server Error")
    }

    return nil
}