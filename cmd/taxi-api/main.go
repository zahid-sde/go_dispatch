package main

import ( 
	"fmt"
	"net/http"
)

func main(){

	http.HandleFunc("/health", healthHandler)

	fmt.Println("Taxi server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Taxi service is healthy!"))
}

