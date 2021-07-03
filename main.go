package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vitorvidaldev/MSN/user/domain"
)

func main() {
	r := mux.NewRouter()
	domain.Router(r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
