package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not authenticated")
)

// Login handler for access
func (s *server) SignIn() http.HandlerFunc {

	type request struct {
		Email string    `json:"email"`
		Password string `json:"password"`
	}
	
	type JWT struct {
		Token string `json:"token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err:= json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		user, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !user.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		token, err := CreateJWT(user.Id)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusOK, &JWT{Token:token})
	}
}



