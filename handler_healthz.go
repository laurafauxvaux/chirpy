package main

import "net/http"

func (_ *apiConfig) handlerHealthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
