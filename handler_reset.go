package main

import "net/http"

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, req *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}

	ctx := req.Context()

	cfg.fileserverHits.Swap(0)

	if err := cfg.dbQueries.DeleteAllUsers(ctx); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error deleting users database")
		return
	}
}
