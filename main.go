package main

import (
	"fmt"
	"net/http"

	"github.com/vitorvidaldev/MSN/config"
)

func main() {
	config.MongoConfig()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
