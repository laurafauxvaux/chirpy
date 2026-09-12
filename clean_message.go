package main

import (
	"slices"
	"strings"
)

func cleanMessage(msg string) string {
	profaneWords := []string{"kerfuffle", "sharbert", "fornax"}
	message := strings.Split(msg, " ")
	result := []string{}
	for _, word := range message {
		if slices.Contains(profaneWords, strings.ToLower(word)) {
			result = append(result, "****")
			continue
		}
		result = append(result, word)
	}
	return strings.Join(result, " ")
}
