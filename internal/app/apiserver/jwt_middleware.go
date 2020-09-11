package apiserver

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

// AuthJwtVerify verify token and add userID to the request context
func (s *server) AuthJwtVerify(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("Authorization")
		header = strings.TrimSpace(header)

		token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			s.error(w, r, http.StatusForbidden, err)
			return
		}
		claims, _ := token.Claims.(jwt.MapClaims)

		ctx := context.WithValue(r.Context(), "userID", claims["userID"]) // adding the user ID to the context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
