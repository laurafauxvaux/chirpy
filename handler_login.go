package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/laurafauxvaux/chirpy/internal/auth"
	"github.com/laurafauxvaux/chirpy/internal/database"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	ctx := req.Context()

	params := parameters{}

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong")
		return
	}

	user, err := cfg.dbQueries.GetUserByEmail(ctx, params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password")
		return
	}

	match, err := auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil || !match {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password")
		return
	}

	accessToken, err := auth.MakeJWT(user.ID, cfg.secret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	refreshToken := auth.MakeRefreshToken()

	if err := cfg.dbQueries.SaveRefreshToken(ctx, database.SaveRefreshTokenParams{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().AddDate(0, 0, 60),
	}); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	newUser := User{
		ID:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Email:        user.Email,
		Token:        accessToken,
		RefreshToken: refreshToken,
	}

	respondWithJSON(w, http.StatusOK, newUser)
}
