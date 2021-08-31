package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode("Health")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Teste")
}
