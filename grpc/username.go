package grpc

import (
	"context"
	"errors"
	"github.com/pepeunlimited/microservice-kit/headers"
)

var (
	ErrMissingXJwtUsername 		= errors.New("headers: missing X-JWT-Username")
)

func UsernameDecode(ctx context.Context) (string, error) {
	username, valid := decode(headers.XJwtUsername, ctx)
	if !valid {
		return "", ErrMissingXJwtUsername
	}
	return username, nil
}

func GetUsername(username string) context.Context {
	return add(username, headers.XJwtUsername)
}