package apiserver

import (
	"net/http"
)

// TokenAuthMiddleware verify token
func (s *server) TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := TokenValid(r)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
