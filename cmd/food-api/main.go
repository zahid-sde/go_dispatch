package main

import (
	"fmt"
	"net/http"
)
	
func main() {

	http.HandleFunc("/healths", healthHandler)
    
	fmt.Println("Food server starting on port 8081...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}


func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Food service is healthy!"))
}