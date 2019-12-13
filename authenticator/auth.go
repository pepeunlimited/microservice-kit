package server

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/bcrypt"
)


const (
	MediumBytes int = 32
)

var (
	ErrCryptoExp = errors.New("crypto: password reset has been expired")
	ErrCryptoTicket = errors.New("crypto: ticket database returned error")
)

type Authenticator interface {
	//hash the password/token using bcrypt (default cost)
	Hash(password string) (string, error)

	//check (validate) the hashed password/token (bcrypt)
	Check(hashed string, password string) error

	Random() (string, error)
}

type Crypto struct {}

func (c Crypto) Random() (string, error) {
	return c.token(MediumBytes)
}

func (Crypto) token(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), err
}

func (Crypto) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (Crypto) Check(hashed string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}

func NewAuthenticator() Authenticator {
	return Crypto{}
}