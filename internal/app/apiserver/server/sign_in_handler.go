package server

import (
	"encoding/json"
	"errors"
	"net/http"
	
	"github.com/zlobste/mint-rest-api/internal/app/apiserver/helpers"
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
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		user, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !user.ComparePassword(req.Password) {
			helpers.Error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		token, err := helpers.CreateJWT(user.Id)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		
		helpers.Respond(w, r, http.StatusOK, &JWT{Token:token})
	}
}



