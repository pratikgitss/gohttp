package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func readinessHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (cfg *apiConfig) handlerReqCounts(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load())))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, req)
	})
}

func (cfg *apiConfig) handlerResetCount(w http.ResponseWriter, req *http.Request) {
	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
}

func main() {
	const port = "8080"
	const filepath = "."
	mux := http.NewServeMux()

	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}
	cfg := apiConfig{}

	fs := http.FileServer(http.Dir(filepath))

	mux.Handle("/app/", cfg.middlewareMetricsInc(http.StripPrefix("/app/", fs)))
	mux.HandleFunc("/healthz", readinessHandler)

	mux.HandleFunc("/metrics", cfg.handlerReqCounts)
	mux.HandleFunc("/reset", cfg.handlerResetCount)
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())

}
