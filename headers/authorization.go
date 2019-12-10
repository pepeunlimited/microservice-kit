package headers

import (
	"errors"
	"net/http"
	"strings"
)

const (
	Authorization 			= "Authorization"
	BearerPrefix 			= "Bearer "
)

var (
	ErrMissingAuthorizationBearer 		= errors.New("headers: missing authorization bearer")
)

func SetBearer(token string, r *http.Request) {
	r.Header.Set(Authorization, BearerPrefix+token)
}

func GetBearer(bearer string) (string, error) {
	if !strings.HasPrefix(bearer, BearerPrefix) {
		return "", ErrMissingAuthorizationBearer
	}
	return strings.TrimPrefix(bearer, BearerPrefix), nil
}
