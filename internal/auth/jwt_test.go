package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJWTValid(t *testing.T) {
	userID := uuid.New()
	secret := "my-secret"

	token, err := MakeJWT(userID, secret, time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT returned an error: %v", err)
	}

	ID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("ValidateJWT returned an error: %v", err)
	}

	if ID != userID {
		t.Errorf("got %v, want %v", ID, userID)
	}
}

func TestJWTExpired(t *testing.T) {
	userID := uuid.New()
	secret := "my-secret"

	token, err := MakeJWT(userID, secret, -time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT returned an error: %v", err)
	}

	_, err = ValidateJWT(token, secret)
	if err == nil {
		t.Fatalf("ValidateJWT should return an error: token expired")
	}
}

func TestJWTWrongSecret(t *testing.T) {
	userID := uuid.New()
	secret := "my-secret"

	token, err := MakeJWT(userID, secret, time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT returned an error: %v", err)
	}

	_, err = ValidateJWT(token, "other-secret")
	if err == nil {
		t.Fatalf("ValidateJWT should return an error: wrong secret")
	}
}

func TestGetBearerToken(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Authorization", "Bearer token-string")

	token, err := GetBearerToken(headers)

	if err != nil {
		t.Fatalf("GetBearerToken returned an error: %v", err)
	}

	if token != "token-string" {
		t.Fatalf("token wasn't extracted properly")
	}
}

func TestGetBearerTokenNoAuthHeader(t *testing.T) {
	headers := make(http.Header)

	_, err := GetBearerToken(headers)

	if err == nil {
		t.Fatalf("GetBearerToken should return an error: authorization header doesn't exist")
	}

}

func TestGetBearerTokenNoBearer(t *testing.T) {
	headers := make(http.Header)
	headers.Set("Authorization", "NoBearer no-token")

	_, err := GetBearerToken(headers)

	if err == nil {
		t.Fatalf("GetBearerToken should return an error: no bearer token in authorization header")
	}

}
