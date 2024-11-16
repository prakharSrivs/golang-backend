package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prakharsrivs/kirana-club-assignment/controllers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/submit", controllers.JobSubmissionController).Methods("POST")
	router.HandleFunc("/api/status", controllers.JobInfoController).Methods("GET")
	log.Println("Server running on PORT 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
