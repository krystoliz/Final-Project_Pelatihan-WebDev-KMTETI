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
		// check if theres a title query parameter
		title := r.URL.Query().Get("title")

		if title != "" {
			//Get specific book by name
			book, err := service.GetBookByTitle(title)
			if err != nil {
				if err.Error()== "Book bot found"{
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}



		//Get all book if no title parameter
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
		defer r.Body.Close()
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Book has been created successfully")
	case "PUT":
		err := service.UpdateBook(r.Body)
        if err != nil {
            switch err.Error() {
            case "Bad Request", "Title is required":
                http.Error(w, err.Error(), http.StatusBadRequest)
            case "Book not found":
                http.Error(w, err.Error(), http.StatusNotFound)
            default:
                http.Error(w, "Internal server error", http.StatusInternalServerError)
            }
            return
        }
        defer r.Body.Close()
        w.Header().Add("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode("Book has been updated successfully")
        return
	case "DELETE":
		title := r.URL.Query().Get("title")
        err := service.DeleteBook(title)
        if err != nil {
            switch err.Error() {
            case "Title is required":
                http.Error(w, err.Error(), http.StatusBadRequest)
            case "Book not found":
                http.Error(w, err.Error(), http.StatusNotFound)
            default:
                http.Error(w, "Internal server error", http.StatusInternalServerError)
            }
            return
        }
        w.Header().Add("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode("Book has been deleted successfully")
        return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
}
	w.Write([]byte("Hello from book"))
}