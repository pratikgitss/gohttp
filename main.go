package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	const filepath = "."
	mux := http.NewServeMux()

	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	mux.Handle("/", http.FileServer(http.Dir(filepath)))
	mux.Handle("/assets/logo.png", http.FileServer(http.Dir(filepath)))

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())

}
