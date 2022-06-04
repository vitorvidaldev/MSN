package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vitorvidaldev/msn/src/application/controller"
	"github.com/vitorvidaldev/msn/src/config"
)

func main() {
	db := config.InitPGConnection()
	defer db.Close()

	r := mux.NewRouter()
	subR := r.PathPrefix("/rest/v1").Subrouter()

	// users
	subR.HandleFunc("/users", controller.CreateUser).Methods("POST")
	subR.HandleFunc("/users", controller.LoginUser).Methods("POST")
	subR.HandleFunc("/users/{userId}", controller.GetUser).Methods("GET")
	subR.HandleFunc("/users/{userId}", controller.DeleteUser).Methods("DELETE")

	// conversations
	subR.HandleFunc("/conversations", controller.StartConversation).Methods("POST")
	subR.HandleFunc("/conversations/{conversationId}", controller.DeleteConversation).Methods("DELETE")

	// messages
	subR.HandleFunc("/conversations/{conversationId}/messages/{messageId}", controller.SendMessage).Methods("DELETE")
	subR.HandleFunc("/messages/history/{userId}", controller.GetMessageHistory).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
