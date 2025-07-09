package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	mux := http.NewServeMux() // serve mux is not a multiplexor per-se but a traffic director
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/healthz", handlerReadiness)

	fmt.Printf("serving on port %s\n", port)
	log.Fatal(server.ListenAndServe())

	return
}
