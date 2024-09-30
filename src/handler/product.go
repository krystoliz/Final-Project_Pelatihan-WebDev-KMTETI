package handler

import "net/http"

func ProductHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed) 
		w.Write([]byte("Method Not Allowed"))
		return
	}
	
    w.Write([]byte("Hello from /api/product"))
}