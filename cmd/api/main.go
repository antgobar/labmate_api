package main

import (
	"log"
	"os"
)

func main() {

	port := "8080"

	var addr string
	if os.Getenv("LABMATE_API_ENV") == "DEV" {
		addr = "localhost:" + port
	} else {
		addr = ":" + port
	}

	cfg := config{
		addr: addr,
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
