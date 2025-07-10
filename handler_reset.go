package main

import "net/http"

func (apiCfg *apiConfig) handlerResetMetrics(w http.ResponseWriter, r *http.Request) {
	apiCfg.fileserverHits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	return
}
