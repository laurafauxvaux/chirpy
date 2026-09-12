package main

import (
	"encoding/json"
	"net/http"
)

func handlerValidateChirp(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	params := parameters{}

	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong")
		return
	}
	if len(params.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	cleanedBody := cleanMessage(params.Body)

	respondWithJSON(w, http.StatusOK, struct {
		CleanedBody string `json:"cleaned_body"`
	}{
		CleanedBody: cleanedBody,
	})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	resp := struct {
		Error string `json:"error"`
	}{
		Error: msg,
	}

	respondWithJSON(w, code, resp)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
