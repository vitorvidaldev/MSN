package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(`{"alive": true}`)
	if err != nil {
		log.Fatal(err)
	}
}
