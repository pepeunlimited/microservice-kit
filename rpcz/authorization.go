package rpcz

import (
	"context"
	"errors"
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/jwt"
)

var (
	ErrMissingAuthorization 	= errors.New("headers: missing Authorization")
)

func AddAuthorization(token string) context.Context {
	return add("Bearer "+token, headers.Authorization)
}

func GetAuthorization(ctx context.Context) (string, error) {
	auth, valid := decode(headers.Authorization, ctx)
	if !valid {
		return "", ErrMissingAuthorization
	}
	return auth, nil
}

func GetAuthorizationWithoutPrefix(ctx context.Context) (string, error) {
	authorization, err := GetAuthorization(ctx)
	if err != nil {
		return "", err
	}
	return jwt.GetBearer(authorization)
}
