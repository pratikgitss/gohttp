package main

import (
	"log"
	"net/http"
)

func readinessHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	const port = "8080"
	const filepath = "."
	mux := http.NewServeMux()

	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	fs := http.FileServer(http.Dir(filepath))

	mux.Handle("/app/", http.StripPrefix("/app/", fs))

	mux.HandleFunc("/healthz", readinessHandler)

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())

}
