package apiserver

import (
	"net/http"
)

func (s *server) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenAuth, err := ExtractTokenMetadata(r)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}
		user, err := s.store.User().FindById(tokenAuth.UserId)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}
		user.Sanitize()
		s.respond(w, r, http.StatusCreated, user)
	}
}
