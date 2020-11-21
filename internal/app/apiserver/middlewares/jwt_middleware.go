package middlewares

import (
	"net/http"
	
	"github.com/zlobste/mint-rest-api/internal/app/apiserver/helpers"
)

// TokenAuthMiddleware verify token
func TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := helpers.TokenValid(r)
		if err != nil {
			helpers.Error(w, r, http.StatusUnauthorized, err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
