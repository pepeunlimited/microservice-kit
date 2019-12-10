package headers

import (
	"context"
	"errors"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/siimooo/pepeutil/nginxjwt"
	"strconv"
)

const (
		XJwtUserId 	 				= "X-JWT-UserId"
)

var (
		ErrMissingXJwtUserId 		= errors.New("headers: missing X-JWT-UserId header")
		ErrNotValidXJwtUserId 		= errors.New("headers: X-JWT-UserId is not int64")
)

func UserIdDecode(ctx context.Context) (int64, error) {
	userId, valid := decode(nginxjwt.XJwtUserId, ctx)
	if !valid {
		return 0, ErrMissingXJwtUserId
	}
	parsed, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return 0, ErrNotValidXJwtUserId
	}
	return parsed, nil
}

func UserId() middleware.Middleware {
	return middlewarez(XJwtUserId)
}
