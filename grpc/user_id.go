package grpc

import (
	"context"
	"errors"
	"github.com/pepeunlimited/microservice-kit/headers"
	"strconv"
)

var (
	ErrMissingXJwtUserId 		= errors.New("headers: missing X-JWT-UserId")
	ErrNotValidXJwtUserId 		= errors.New("headers: X-JWT-UserId is not int64")
)

func AddUserId(userId int64) context.Context {
	return add(strconv.FormatInt(userId, 10), headers.XJwtUserId)
}

func GetUserId(ctx context.Context) (int64, error) {
	userId, valid := decode(headers.XJwtUserId, ctx)
	if !valid {
		return 0, ErrMissingXJwtUserId
	}
	parsed, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return 0, ErrNotValidXJwtUserId
	}
	return parsed, nil
}