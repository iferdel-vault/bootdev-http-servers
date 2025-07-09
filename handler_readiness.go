package main

import (
	"net/http"
)

// // The HandlerFunc type is an adapter to allow the use of
// // ordinary functions as HTTP handlers. If f is a function
// // with the appropriate signature, HandlerFunc(f) is a
// // [Handler] that calls f.
// type HandlerFunc func(ResponseWriter, *Request)
//
// // ServeHTTP calls f(w, r).
// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
// 	f(w, r)
// }

// qué pasa aquí... Para poder servir HTTP responses, se requiere de una función
// que haga algo con un request o bien con un response, que pueda operar con esos
// componentes de http.
//

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
	return
}
