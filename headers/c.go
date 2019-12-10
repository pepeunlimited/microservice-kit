package headers

import (
	"context"
	"github.com/pepeunlimited/microservice-kit/middleware"
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