package rpcz

import (
	"context"
	"errors"
	"github.com/pepeunlimited/microservice-kit/headers"
	"github.com/pepeunlimited/microservice-kit/validator"
	"strings"
)

var (
	ErrMissingRoles 			= errors.New("headers: missing X-JWT-Roles")
)

func AddRoles(roles []string) context.Context {
	if roles == nil || len(roles) == 0 {
		return nil
	}
	var str string
	for i, role := range roles {
		if len(roles) - 1 > i {
			str += role+","
		} else {
			str += role
		}

	}
	return add(str, headers.XJwtRoles)
}

func AddRole(role string) context.Context {
	return AddRoles([]string{role})
}

func GetRoles(ctx context.Context) ([]string, error) {
	role, valid := decode(headers.XJwtRoles, ctx)
	if !valid {
		return nil, ErrMissingRoles
	}
	if validator.IsEmpty(role) {
		return []string{}, nil
	}
	roles := strings.Split(role, ",")
	return roles, nil
}