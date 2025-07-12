package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/iferdel-vault/bootdev-http-servers/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	platform       string
	fileserverHits atomic.Int32
	db             *database.Queries
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
	const filepathRoot = "."
	const port = "8080"

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error reading .env file: %v", err)
	}

	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM must be set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error opening connection to db: %v", err)
	}
	dbQueries := database.New(dbConn)

	apiCfg := apiConfig{
		platform:       platform,
		fileserverHits: atomic.Int32{},
		db:             dbQueries,
	}

	mux := http.NewServeMux() // serve mux is not a multiplexor per-se but a traffic director
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.Handle("/app/", http.StripPrefix("/app", apiCfg.middlewareMetricsInc(http.FileServer(http.Dir(filepathRoot)))))
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerResetUsers)
	mux.HandleFunc("POST /api/validate_chirp", handlerValidateChirp)

	fmt.Printf("serving on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
