package main

import (
	"fmt"
	"net/http"
)

func (apiCfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	metrics := apiCfg.fileserverHits.Load()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
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
