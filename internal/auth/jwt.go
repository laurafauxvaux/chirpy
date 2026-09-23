package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	now := time.Now().UTC()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "chirpy-access",
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate((now.Add(expiresIn))),
		Subject:   userID.String(),
	})

	signedJWT, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}

	return signedJWT, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	claimsValue := jwt.RegisteredClaims{}
	claims := &claimsValue

	token, err := jwt.ParseWithClaims(tokenString, claims, func(*jwt.Token) (any, error) { return []byte(tokenSecret), nil })
	if err != nil {
		return uuid.Nil, err
	}

	userID, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return uuid.Nil, err
	}

	return userUUID, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	authData := headers.Get("Authorization")
	if authData == "" {
		return "", fmt.Errorf("authorization header doesn't exist")
	}

	if !strings.HasPrefix(authData, "Bearer ") {
		return "", fmt.Errorf("no bearer token in authorization header")
	}

	tokenString := strings.TrimSpace(strings.TrimPrefix(authData, "Bearer "))

	return tokenString, nil
}
