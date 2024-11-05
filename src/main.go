package main

import (
	"fmt"
	"net/http"
	/* "github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/service" */
	/* "github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/db" */
	 "github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/controller"
	h1 "github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/api"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from /api/users"))
}

func main() {

	h := http.NewServeMux()
	
	s := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from server"))
	})

	h.HandleFunc("/api/user", userHandler)

	h.HandleFunc("/api/books", controller.BookController) //if p is lowercase it wont be accessible to this file. If p is uppercase it will be accessible.
	h.HandleFunc("/api/books-serverless", h1.BookHandler) 
	h.HandleFunc("/api/test-db", controller.NewProductHandler)

	fmt.Println("HTTP Server is running on port 8080")
	err := s.ListenAndServe()
	if err != nil { //error handling
		fmt.Println(err.Error())
	}

}
