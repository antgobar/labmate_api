package main

import (
	"fmt"
	"net/http"
)

func main() {
	api := &api{addr: "localhost:8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)
	fmt.Println("Running server ...")
	srv.ListenAndServe()
}
