package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var addr string
	if host := os.Getenv("LABMATE_API_HOST"); host != "" {
		addr = fmt.Sprintf("%s:8080", host)
	} else {
		addr = ":8080"
	}

	api := &api{addr: addr}

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
