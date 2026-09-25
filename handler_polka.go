package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/laurafauxvaux/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerUpgradeUser(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	type polkaWebhooksParams struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		} `json:"data"`
	}
	params := polkaWebhooksParams{}

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	apiKey, err := auth.GetAPIKey(req.Header)
	if err != nil || apiKey != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if _, err := cfg.dbQueries.UpgradeUser(ctx, params.Data.UserID); err != nil {
		respondWithError(w, http.StatusNotFound, "user can't be found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
