package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var addr string
	if env := os.Getenv("ENV"); env == "DEV" {
		addr = "localhost:8080"
	} else {
		addr = ":8080"
	}

	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)
	fmt.Printf("Running server on addr %s\n", addr)
	srv.ListenAndServe()
}
