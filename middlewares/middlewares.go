package middlewares

import "net/http"

func SetMiddlewareJSON(_next http.HandlerFunc) http.HandlerFunc {
	return func(_w http.ResponseWriter, _r *http.Request) {
		_w.Header().Set("Content-Type", "application/json")
		_next(_w, _r)
	}
}
