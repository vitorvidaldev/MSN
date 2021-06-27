package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitorvidaldev/MSN/application/controller"
)

func main() {
	// Router
	r := mux.NewRouter()

	// Routes
	s := r.PathPrefix("/user").Subrouter()
	s.HandleFunc("/", controller.GetUsers).Methods("GET")
	s.HandleFunc("/{id}", controller.GetUserById).Methods("GET")
	s.HandleFunc("/", controller.CreateUser).Methods("POST")
	s.HandleFunc("/", controller.UpdateUser).Methods("PUT")
	s.HandleFunc("/{id}", controller.DeleteUser).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
