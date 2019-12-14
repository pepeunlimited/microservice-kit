package rpcz

import (
	"context"
	"errors"
	"github.com/pepeunlimited/microservice-kit/headers"
)

var (
	ErrMissingXForwardedFor 	= errors.New("headers: missing X-Forwarded-For")
)

func XForwardedForDecode(ctx context.Context) (string, error) {
	xforwardedfor, valid := decode(headers.XForwardedFor, ctx)
	if !valid {
		return "", ErrMissingXForwardedFor
	}
	return xforwardedfor, nil
}

func GetXForwardedFor(ip string) context.Context {
	return add(ip, headers.XForwardedFor)
}
