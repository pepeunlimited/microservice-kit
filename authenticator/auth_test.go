package server

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCrypto_HashPw(t *testing.T) {
	password := "password"
	authenticator := NewAuthenticator()
	hash, err := authenticator.Hash(password)
	if err != nil {
		t.Fatal(err)
	}
	if len(hash) == 0 {
		t.Fatal("should not be zero")
	}
	if err := authenticator.Check(hash, password); err != nil {
		t.Fatal(err)
	}
}

func TestCrypto_HashPwNotSame(t *testing.T) {
	password1 := "password"
	password2 := "passu"
	authenticator := NewAuthenticator()
	hash, err := authenticator.Hash(password1)
	if err != nil {
		t.Fatal(err)
	}
	if len(hash) == 0 {
		t.Fatal("should not be zero")
	}
	if err := authenticator.Check(hash, password2); err != bcrypt.ErrMismatchedHashAndPassword {
		t.Fatal("should be ErrMismatchedHashAndPassword")
	}
}


func TestCrypto_Random(t *testing.T) {
	authenticator := NewAuthenticator()
	token, err := authenticator.Random()
	if err != nil {
		t.Fatal(err)
	}
	if len(token) != 44 {
		t.Fatal("should be 44")
	}
}