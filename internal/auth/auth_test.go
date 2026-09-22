package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestAuthWithMatch(t *testing.T) {
	password := "pa$$word"
	hash, err := HashPassword(password)

	if err != nil {
		t.Fatalf("HashPassword returned an error: %v", err)
	}

	if password == hash {
		t.Errorf("expected password to be hashed")
	}

	res, err := CheckPasswordHash(password, hash)

	if err != nil {
		t.Fatalf("CheckPasswordHash returned an error: %v", err)
	}

	if !res {
		t.Errorf("hash doesn't match")
	}
}

func TestAuthNoMatch(t *testing.T) {
	wrongPassword := "pa$$word"
	hash, err := HashPassword("passw0rd")

	if err != nil {
		t.Fatalf("HashPassword returned an error: %v", err)
	}

	res, err := CheckPasswordHash(wrongPassword, hash)

	if err != nil {
		t.Fatalf("CheckPasswordHash returned an error: %v", err)
	}

	if res {
		t.Errorf("hash shouldn't match")
	}
}

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
