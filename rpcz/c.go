package rpcz

import (
	"context"
	"github.com/pepeunlimited/microservice-kit/validator"
)

func add(value string, header string) context.Context {
	return context.WithValue(context.TODO(), header, value)
}

func decode(header string, ctx context.Context) (string, bool)  {
	value := ctx.Value(header)
	if value == nil {
		return "", false
	}
	str := value.(string)
	if validator.IsEmpty(str) {
		return "", false
	}
	return str, true
}