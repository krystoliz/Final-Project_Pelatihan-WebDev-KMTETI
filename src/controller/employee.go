package controller

import (
	/* "context" */
	"encoding/json"
	/* "fmt" */
	"net/http"
	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/service"
	/* "github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/db"
	"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" */
)

func EmployeeController(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		// check if theres a title query parameter
		/* name := r.URL.Query().Get("name")

		if name != "" {
			//Get specific book by name
			employee, err := service.GetEmployeeByName(name)
			if err != nil {
				if err.Error()== "Employee not found"{
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(employee)
			return
		} */

		//Get all book if no title parameter
		/* data, err := service.GetAllEmployee()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data) */
		/* w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bookList) */
		/* return */

	case "POST":
		err := service.CreateEmployee(r.Body)
		if err != nil {
			if err.Error() == "bad request"{
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Employee created")
		

	case "PUT":
		/* err := service.UpdateEmployee(r.Body)
		if err != nil {
			if err.Error() == "bad request"{
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Employee updated"))
		return */

	case "DELETE":
		/* err := service.DeleteEmployee(r.Body)
		if err != nil {
			if err.Error() == "bad request"{
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Employee deleted"))
		return */

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}