package main

import (
	"net/http"

	"github.com/laurafauxvaux/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	token, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := cfg.dbQueries.RevokeRefreshToken(ctx, token); err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
