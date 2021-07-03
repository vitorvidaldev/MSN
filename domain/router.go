package domain

import (
	"github.com/gorilla/mux"
	"github.com/vitorvidaldev/MSN/application/controller"
)

func Router(r *mux.Router) {
	userRouter(r)
	healthcheckRouter(r)
}

func healthcheckRouter(r *mux.Router) {
	s := r.PathPrefix("/health").Subrouter()
	s.HandleFunc("/", controller.HealthCheck).Methods("GET")
}

func userRouter(r *mux.Router) {
	s := r.PathPrefix("/user").Subrouter()
	s.HandleFunc("/", controller.GetUsers).Methods("GET")
	s.HandleFunc("/{id}", controller.GetUserById).Methods("GET")
	s.HandleFunc("/", controller.CreateUser).Methods("POST")
	s.HandleFunc("/{id}", controller.UpdateUser).Methods("PUT")
	s.HandleFunc("/{id}", controller.DeleteUser).Methods("DELETE")
}
