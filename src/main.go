package main

import ("fmt"
		"net/http")
func main(){
	s := &http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello from server"))
	})
	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello from /api/user"))
	})
	http.HandleFunc("/api/product", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello from /api/product"))
	})
	fmt.Println("HTTP Server is running on port 8080")
	err := s.ListenAndServe()
	if err != nil { //error handling
		fmt.Println(err.Error())
	}
	
}