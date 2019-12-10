package headers

import (
	"context"
	"errors"
	"github.com/pepeunlimited/microservice-kit/middleware"
)

const (
	XForwardedFor			= "X-Forwarded-For"
)

var (
	ErrMissingXForwardedFor 	= errors.New("headers: missing X-Forwarded-For header")
)

func XForwardedForDecode(ctx context.Context) (string, error) {
	xforwardedfor, valid := decode(XForwardedFor, ctx)
	if !valid {
		return "", ErrMissingXForwardedFor
	}
	return xforwardedfor, nil
}

func XForwardedForz() middleware.Middleware {
	return middlewarez(XForwardedFor)
}