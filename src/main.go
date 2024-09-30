package main

import ("fmt"
		"net/http")

func userHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from /api/users"))
}

func productHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Method)
	if r.Method == "GET" {
		fmt.Println("GET Method retrieved")
	}
	if r.Method == "POST" {
        fmt.Println("POST Method retrieved")
    }
    w.Write([]byte("Hello from /api/product"))
}

 
func main(){
	s := &http.Server{
		Addr: ":8080",
	}


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello from server"))
	})

	http.HandleFunc("/api/user", userHandler)
	http.HandleFunc("/api/product", productHandler)

	fmt.Println("HTTP Server is running on port 8080")
	err := s.ListenAndServe()
	if err != nil { //error handling
		fmt.Println(err.Error())
	}
	
}