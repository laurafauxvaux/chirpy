package main

import "net/http"

func (cfg *apiConfig) handlerReset(_ http.ResponseWriter, _ *http.Request) {
	cfg.fileserverHits.Swap(0)
}
