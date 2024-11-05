# Final-Project_Pelatihan-WebDev-KMTETI

# Book & Employee Management System

A RESTful API service built with Go and MongoDB for managing book data and employee data. This project was developed as part of the Final Project for KMTETI Web Development Training.

## Features
- View book list
- View the detail of a specific book
- Add a new book to database
- Make an update to the stock and price of a specific book
- Create new employee records
- View employee list with basic information
- Data validation for employee creation
- MongoDB integration for data persistence

## Tech Stack

- **Backend:** Go (Golang)
- **Database:** MongoDB
- **Packages:**
  - `go.mongodb.org/mongo-driver` - MongoDB driver for Go
  - `encoding/json` - JSON encoding/decoding
  - Other standard Go packages

## Prerequisites

Before running this project, make sure you have the following installed:
- Go (version 1.16 or higher)
- MongoDB (version 4.0 or higher)
- Git

## Installation

1. Clone the repository
```bash
git clone https://github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI.git
cd Final-Project_Pelatihan-WebDev-KMTETI
```

2. Install dependencies
```bash
go mod download
```

3. Set up your MongoDB connection
- Create a MongoDB database
- Update the connection string in the project configuration (if necessary)

4. Run the application
```bash
go run main.go
```

## API Endpoints
### Get Books List
- **Endpoint:** `GET localhost:8080/api/books` or  `GET https://final-project-pelatihan-web-dev-kmteti.vercel.app/api/book `
- **Description:** Retrieves a list of all books with basic information
- **Response:**
```json
{
    "data": [
        {
            "title": "Unf/Air",
            "author": "ARTMS",
            "price": 250000
        }
    ]
}
```

### Get Books List
- **Endpoint:** `GET localhost:8080/api/books?title=Confessions` or  `GET https://final-project-pelatihan-web-dev-kmteti.vercel.app/api/book?Unf/Air `
- **Description:** Retrieves detailed information of a specific book
- **Response:**
```json
{
    "title": "Unf/Air",
    "author": "ARTMS",
    "stock": 10,
    "year_released": 2024,
    "price": 250000
}
```
### Create a book
- **Endpoint:** `POST localhost:8080/api/book` or `POST https://final-project-pelatihan-web-dev-kmteti.vercel.app/api/book`
- **Description:** Creates a new book record
- **Request Body:**
```json
{
        "title": "Secret Diary",
        "Author": "Loossemble",
        "Stock": 330,
        "Year_released": 2024,
        "Price": 120001
}
```
- **Required Fields:** title

### Update a book
- **Endpoint:** `PUT localhost:8080/api/books` or  `PUT https://final-project-pelatihan-web-dev-kmteti.vercel.app/api/book`
- **Description:** Update a specific book record
- **Request Body:**
```json
{
    "title":"Confessions",
    "price": 250000,
    "stock": 10
}
```
- **Required Fields:** title, price, stock

### Delete a book
- **Endpoint:** `DELETE localhost:8080/api/books?title=Unf/Air` or `DELETE https://final-project-pelatihan-web-dev-kmteti.vercel.app/api/book?title=Unf/Air`
- **Description:** Delete a specific book record
- **Response:**
```json
"Book has been deleted successfully"
```
- **Required Fields:** title

### Create Employee
- **Endpoint:** `POST localhost:8080/api/employee` OR `POST https://final-project-pelatihan-web-dev-kmteti.vercel.app/api/employee`
- **Description:** Creates a new employee record
- **Request Body:**
```json
{
    "name":"YeoJin",
    "nik":"123456",
    "educationLevel":"SMA",
    "dateStarted":"2022-09-22",
    "workStatus":"Full-time"

}
```
- **Required Fields:** name, nik, educationLevel, dateStarted, workStatus
- **Work Status Options:** "Contract" or "Full-time"

### Get Employees List
- **Endpoint:** `GET localhost:8080/api/employee` or `GET https://final-project-pelatihan-web-dev-kmteti.vercel.app/api/employee`
- **Description:** Retrieves a list of all employees with basic information
- **Response:**
```json
[
    [
    {
        "name": "Hyeju",
        "dateStarted": "2022-09-22T00:00:00Z",
        "workStatus": "Full-time"
    },
    {
        "name": "Go Won",
        "dateStarted": "2022-09-22T00:00:00Z",
        "workStatus": "Full-time"
    },
    {
        "name": "YeoJin",
        "dateStarted": "2022-09-22T00:00:00Z",
        "workStatus": "Full-time"
    }
]
]
```

## Project Structure

```
.
├── api/
│   ├── book.go
│   └── employee.go
├── main.go
├── src/
│   ├── db/
│   │   └── db.go
│   ├── controller/
│   │   ├── book.go
│   │   └── employee.go
│   ├── model/
│   │   ├── book.model.go
│   │   └── employee.model.go
│   └── service/
│       ├── book.model.go
│       └── employee.model.go
├── go.mod
├── go.sum
├── README.md
├── .gitignore
└── dist/
    └── server.exe
```

## Data Models

### Employee
```go
type Employee struct {
    ID              primitive.ObjectID
    Name            string
    NIK             string
    Education_level string
    Date_started    time.Time
    Work_status     WorkStatus
}
```

```go
type Book struct {
	ID            primitive.ObjectID 
	Title         string             
	Author        string             
	Stock         int                
	Year_released int                
	Price         int                
}

```

## Error Handling

The API returns appropriate HTTP status codes:
- 200: Successful operation
- 400: Bad Request (invalid input)
- 405: Method Not Allowed
- 500: Internal Server Error

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Authors

- **Jhon Samuel Kudadiri** - *Initial work* - [krystoliz](https://github.com/krystoliz)
- **Claude + Copilot** - *API Endpoints and Service Function Code* 
(https://claude.ai/chat/abdda77e-3404-4c0c-a14d-e7d2976bdfd4) - 60% GenAI + 40% Meeting 3-6 code

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

- KMTETI Web Development Training Program
- All contributors and participants

## Future Improvements

- Add authentication and authorization
- Implement employee data update and deletion
- Add more detailed employee information
- Implement search and filtering capabilities
- Add unit tests and integration tests
