package handler

import (
	/* "context" */
	"encoding/json"
	/* "fmt" */

	"net/http"

	/* "github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/db"
	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/model" */
	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/service"
	/* "go.mongodb.org/mongo-driver/bson" */
	/* "go.mongodb.org/mongo-driver/bson/primitive" */
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
		err := service.CreateBook(r.Body)
		if err != nil {
			if err.Error() == "bad request"{
				http.Error(w, err.Error(), http.StatusBadRequest)

			}
			http.Error(w, "internal server error" , http.StatusInternalServerError)
			
		}
		
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Book has been created successfully")
	case "PUT":

	case "DELETE":

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
}
	w.Write([]byte("Hello from book"))
}