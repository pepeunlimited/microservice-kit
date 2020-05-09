package cryptoz

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)


const (
	MediumBytes int = 32
)

type Crypto interface {
	//hash the password/token using bcrypt (default cost)
	Hash(password string) (string, error)

	//check (validate) the hashed password/token (bcrypt)
	Check(hashed string, password string) error

	Random() (string, error)

	RandomAt(length int, encode bool) (string, error)
}

type crypto struct {}

func (c crypto) RandomAt(length int, encode bool) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	if !encode {
		return fmt.Sprintf("%X", bytes), nil
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func (c crypto) Random() (string, error) {
	return c.RandomAt(MediumBytes, true)
}

func (crypto) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (crypto) Check(hashed string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}

func NewCrypto() Crypto {
	return crypto{}
}