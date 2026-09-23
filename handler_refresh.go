package main

import (
	"net/http"
	"time"

	"github.com/laurafauxvaux/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	token, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userID, err := cfg.dbQueries.GetUserFromRefreshToken(ctx, token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "token is expired or revoked")
		return
	}

	accessToken, err := auth.MakeJWT(userID, cfg.secret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	resp := struct {
		Token string `json:"token"`
	}{
		Token: accessToken,
	}

	respondWithJSON(w, http.StatusOK, resp)
}
