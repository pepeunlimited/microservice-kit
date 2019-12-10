package headers

import (
	"context"
	"errors"
	"github.com/pepeunlimited/microservice-kit/middleware"
)

const (
	XJwtRole 				= "X-JWT-Role"
)

var (
	ErrMissingRole 			= errors.New("headers: missing X-JWT-Role header")
)


func RoleDecode(ctx context.Context) (string, error) {
	role, valid := decode(XJwtRole, ctx)
	if !valid {
		return "", ErrMissingRole
	}
	return role, nil
}

func Role() middleware.Middleware {
	return middlewarez(XJwtRole)
}