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

	mux.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Printf("serving on port %s\n", port)
	log.Fatal(server.ListenAndServe())

	return
}
