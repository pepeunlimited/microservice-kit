package cryptoz

import (
	"errors"
)

type CryptoMock struct {
	ForgotPasswordAttempts int
	ForgotPasswordCount int
	ForgotToken string
	ForgotTokenSame string
	ForgotPasswordErr bool
}

func (mock CryptoMock) Random() (string, error) {
	if mock.ForgotPasswordErr {
		return "", ErrCryptoMock
	}

	if mock.ForgotPasswordAttempts < mock.ForgotPasswordCount {
		return mock.ForgotTokenSame, nil
	}
	mock.ForgotPasswordCount += 1
	return mock.ForgotToken, nil
}

var (
	ErrCryptoMock = errors.New("crypto_mock: generic error")
)

func (CryptoMock) Hash(password string) (string, error) {
	return password, nil
}

func (CryptoMock) Check(hashed string, password string) error {
	panic("implement me")
}


func NewAuthenticatorMock() Crypto {
	return &CryptoMock{ForgotPasswordAttempts: 5, ForgotPasswordCount:0, ForgotToken: "forgot-token", ForgotTokenSame:"forgot-token-same"}
}



