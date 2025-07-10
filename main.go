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

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	fmt.Println("yeah buddy!!!!!") // this runs on startup, setup of the whole thing
	// the idea is to return a NEW handler that does something and then calls the original 'wrapped' handler (which is called next for sake of semantics)

	// a handlerfunc is a function that converts a function (with the ServeHTTP signature) into a Handler type
	// Handler is an interface that tells how a Handler in HTTP should look like (interface)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func main() {

	apiCfg := apiConfig{}

	const port = "8080"

	mux := http.NewServeMux() // serve mux is not a multiplexor per-se but a traffic director
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.Handle("/app/", http.StripPrefix("/app", apiCfg.middlewareMetricsInc(http.FileServer(http.Dir(".")))))
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerResetMetrics)
	mux.HandleFunc("POST /api/validate_chirp", handlerValidateChirp)

	fmt.Printf("serving on port %s\n", port)
	log.Fatal(server.ListenAndServe())

	return
}
