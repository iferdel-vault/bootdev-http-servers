package main

import (
	"fmt"
	"net/http"
)

func (apiCfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	metrics := apiCfg.fileserverHits.Load()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf(`
		<html>
			<body>
				<h1>Welcome, Chirpy Admin</h1>
				<p>Chirpy has been visited %d times!</p>
			</body>
		</html>
	`, metrics)))
	return
}
