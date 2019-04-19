package middleware

import "net/http"

// Middleware used with compose to generate http handlers
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Compose a set of middleware with a http handler
func Compose(f http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	if len(m) == 0 {
		return f
	}

	return m[0](Compose(f, m[1:cap(m)]...))
}
