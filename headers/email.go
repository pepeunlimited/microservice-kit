package headers

import (
	"context"
	"errors"
	"github.com/pepeunlimited/microservice-kit/middleware"
)

const (
	XJwtEmail    			= "X-JWT-Email"
)

var (
	ErrMissingXJwtEmail 	= errors.New("headers: missing X-JWT-Email header")
)

func EmailDecode(ctx context.Context) (string, error) {
	email, valid := decode(XJwtEmail, ctx)
	if !valid {
		return "", ErrMissingXJwtEmail
	}
	return email, nil
}

func Email() middleware.Middleware {
	return middlewarez(XJwtEmail)
}
