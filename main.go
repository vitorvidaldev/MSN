package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitorvidaldev/MSN/application/controller"
)

// TODO: Fix Error: read ECONNRESET error
func main() {
	// MongoDB initial configuration
	// userCollection := config.MongoConfig()

	// Router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/user", controller.GetUsers).Methods("GET")
	// r.HandleFunc("/user/{id}", controller.GetUserById).Methods("GET")
	// r.HandleFunc("/user", controller.CreateUser).Methods("POST")
	// r.HandleFunc("/user", controller.UpdateUser).Methods("PUT")
	// r.HandleFunc("/user/{id}", controller.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
