package middleware

import "net/http"

//we define a function takes in a http.Handler  and return a http.Handler
type Middleware func( http.Handler) http.Handler

func Adapt(h http.Handler, adapters ...Middleware) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}