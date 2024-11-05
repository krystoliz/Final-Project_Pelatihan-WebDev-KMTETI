package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"time"

	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/db"
	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EmployeeRequest struct {
    Name            string         `json:"name"`
    NIK             string         `json:"nik"`
    EducationLevel  string         `json:"educationLevel"`
    DateStarted     string         `json:"dateStarted"`  // Accepting date as string in format "2006-01-02"
    WorkStatus      model.WorkStatus `json:"workStatus"`
}

type EmployeeListResponse struct {
    Name        string          `json:"name"`
    DateStarted time.Time       `json:"dateStarted"`
    WorkStatus  model.WorkStatus `json:"workStatus"`
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

func GetEmployeesList() ([]EmployeeListResponse, error) {
    // Connect to database
    db, err := db.DBConnection()
    if err != nil {
        return nil, err
    }
    defer db.MongoDB.Client().Disconnect(context.TODO())

    // Get collection
    coll := db.MongoDB.Collection("employee")

    // Define projection to only return specified fields
    projection := bson.D{
        {Key:"name", Value: 1},
        {Key:"date_started",Value: 1},
        {Key: "work_status", Value: 1},
    }

    // Find all employees with projection
    cursor, err := coll.Find(context.TODO(), bson.D{}, &options.FindOptions{
        Projection: projection,
    })
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    // Slice to store results
    var employees []EmployeeListResponse

    // Iterate through cursor
    for cursor.Next(context.TODO()) {
        var emp model.Employee
        if err := cursor.Decode(&emp); err != nil {
            return nil, err
        }
        
        // Create response object with only required fields
        empResponse := EmployeeListResponse{
            Name:        emp.Name,
            DateStarted: emp.Date_started,
            WorkStatus:  emp.Work_status,
        }
        employees = append(employees, empResponse)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return employees, nil
}