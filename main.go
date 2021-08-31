package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitorvidaldev/MSN/application/controller"
)

func main() {
	r := mux.NewRouter()

	s := r.PathPrefix("/user").Subrouter()
	s.HandleFunc("/", controller.GetUsers).Methods("GET")
	s.HandleFunc("/{id}", controller.GetUserById).Methods("GET")
	s.HandleFunc("/", controller.CreateUser).Methods("POST")
	s.HandleFunc("/{id}", controller.UpdateUser).Methods("PUT")
	s.HandleFunc("/{id}", controller.DeleteUser).Methods("DELETE")

	health := r.PathPrefix("/health").Subrouter()
	health.HandleFunc("/", controller.Health).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
