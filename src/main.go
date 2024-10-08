package main

import ("fmt"
		"net/http"
		"github.com/krystoliz/Final-Project_Pelatihan-WebDev-KMTETI/src/handler")

func userHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from /api/users"))
}



 
func main(){
	s := &http.Server{
		Addr: ":8080",
	}
	


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello from server"))
	})

	http.HandleFunc("/api/user", userHandler)
	http.HandleFunc("/api/books", handler.BookHandler) //if p is lowercase it wont be accessible to this file. If p is uppercase it will be accessible.

	fmt.Println("HTTP Server is running on port 8080")
	err := s.ListenAndServe()
	if err != nil { //error handling
		fmt.Println(err.Error())
	}
	
}