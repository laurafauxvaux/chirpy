package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/laurafauxvaux/chirpy/internal/database"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerChirps(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	type params struct {
		Body   string    `json:"body"`
		UserID uuid.UUID `json:"user_id"`
	}

	chirpParams := params{}

	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&chirpParams); err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong")
		return
	}
	if len(chirpParams.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	cleanedBody := cleanMessage(chirpParams.Body)

	chirp, err := cfg.dbQueries.CreateChirp(ctx, database.CreateChirpParams{
		Body:   cleanedBody,
		UserID: chirpParams.UserID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating new chirp")
		return
	}

	newChirp := Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}

	respondWithJSON(w, http.StatusCreated, newChirp)

}

func (cfg *apiConfig) HandlerGetChirps(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	chirps, err := cfg.dbQueries.GetAllChirps(ctx)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error gathering the chirps")
		return
	}

	allChirps := make([]Chirp, len(chirps))

	for i, chirp := range chirps {
		allChirps[i] = Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		}
	}

	respondWithJSON(w, http.StatusOK, allChirps)
}

func (cfg *apiConfig) HandlerGetChirp(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	chirpID := req.PathValue("chirpID")
	chirpUUID, err := uuid.Parse(chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't parse UUID")
		return
	}

	chirp, err := cfg.dbQueries.GetChirp(ctx, chirpUUID)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp can't be found.")
		return
	}

	foundChirp := Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}

	respondWithJSON(w, http.StatusOK, foundChirp)
}
