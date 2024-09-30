package main

import ("fmt"
		"net/http")
func main(){
	s := &http.Server{
		Addr: ":8000",
	}
	
	fmt.Println("HTTP Server is running on port 8000")
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
	
}