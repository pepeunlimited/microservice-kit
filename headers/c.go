package headers

import (
	"context"
	"github.com/pepeunlimited/microservice-kit/middleware"
	"github.com/pepeunlimited/microservice-kit/validator"
	"net/http"
)

func middlewarez(header string) middleware.Middleware {
	return func(h http.Handler) http.Handler {
		return create(h, header)
	}
}

func create(h http.Handler, header string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, header, r.Header.Get(header))
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
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