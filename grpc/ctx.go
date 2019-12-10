package grpc

import (
	"context"
	"github.com/pepeunlimited/microservice-kit/headers"
	"strconv"
)

func add(value string, header string) context.Context {
	return context.WithValue(context.TODO(), header, value)
}

func AddEmail(email string) context.Context {
	return add(email, headers.XJwtEmail)
}

func AddUsername(username string) context.Context {
	return add(username, headers.XJwtUsername)
}

func AddRole(role string) context.Context {
	return add(role, headers.XJwtRole)
}

func AddUserId(userId int64) context.Context {
	return add(strconv.FormatInt(userId, 10), headers.XJwtUserId)
}

func AddXForwaredFor(ip string) context.Context {
	return add(ip, headers.XForwardedFor)
}