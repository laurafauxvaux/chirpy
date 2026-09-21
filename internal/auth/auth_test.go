package auth

import (
	"testing"
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
