package jwt

import (
	"errors"
	"net/http"
	"strings"
)

const (
	BearerPrefix 			= "Bearer "
)

var (
	ErrMissingBearer 		= errors.New("jwt: missing bearer token")
)

func SetBearer(token string, r *http.Request) {
	r.Header.Set("Authorization", BearerPrefix+token)
}

func GetBearer(authorization string) (string, error) {
	if !strings.HasPrefix(authorization, BearerPrefix) {
		return "", ErrMissingBearer
	}
	return strings.TrimPrefix(authorization, BearerPrefix), nil
}