package headers

import (
	"context"
	"errors"
	"github.com/pepeunlimited/microservice-kit/middleware"
)

const (
		XJwtUsername 			= "X-JWT-Username"
)

var (
		ErrMissingXJwtUsername 		= errors.New("headers: missing X-JWT-Username header")
)

func UsernameDecode(ctx context.Context) (string, error) {
	username, valid := decode(XJwtUsername, ctx)
	if !valid {
		return "", ErrMissingXJwtUsername
	}
	return username, nil
}

func Username() middleware.Middleware {
	return middlewarez(XJwtEmail)
}