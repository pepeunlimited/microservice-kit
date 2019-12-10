package grpc

import (
	"context"
	"errors"
	"github.com/pepeunlimited/microservice-kit/headers"
)

var (
	ErrMissingRole 			= errors.New("headers: missing X-JWT-Role")
)

func AddRole(role string) context.Context {
	return add(role, headers.XJwtRole)
}

func GetRole(ctx context.Context) (string, error) {
	role, valid := decode(headers.XJwtRole, ctx)
	if !valid {
		return "", ErrMissingRole
	}
	return role, nil
}