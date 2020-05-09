package cryptoz

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

func TestCrypto_HashPw(t *testing.T) {
	password := "password"
	authenticator := NewCrypto()
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
	authenticator := NewCrypto()
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
	authenticator := NewCrypto()
	token, err := authenticator.Random()
	log.Print(token)
	if err != nil {
		t.Fatal(err)
	}
	if len(token) != 44 {
		t.Fatal("should be 44")
	}
}

func TestCrypto_RandomAt(t *testing.T) {
	at, err := NewCrypto().RandomAt(2, false)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	log.Print(at)
}