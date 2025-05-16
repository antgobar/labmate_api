package main

import (
	"fmt"
	"net/http"
	"os"
)


func main() {
	api := &Api{addr: os.Getenv("API_ADDR")}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)
	fmt.Printf("Running server on addr %s\n", api.addr)
	srv.ListenAndServe()
}
