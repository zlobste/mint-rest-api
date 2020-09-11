package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Login handler for access
func (s *server) SignIn() http.HandlerFunc {

	type request struct {
		Email string    `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err:= json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		user, err := s.store.User().FindByEmail(req.Email)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}

		if !user.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errors.New("Incorrect email or password"))
			return
		}

		token, err := CreateJWT(user.Id)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, token)
	}
}



