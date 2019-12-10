package grpc

import (
	"context"
	"errors"
	"github.com/pepeunlimited/microservice-kit/headers"
)

var (
	ErrMissingXJwtEmail 	= errors.New("headers: missing X-JWT-Email")
)

func AddEmail(email string) context.Context {
	return add(email, headers.XJwtEmail)
}

func GetEmail(ctx context.Context) (string, error) {
	email, valid := decode(headers.XJwtEmail, ctx)
	if !valid {
		return "", ErrMissingXJwtEmail
	}
	return email, nil
}