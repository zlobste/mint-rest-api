package server

import (
	"encoding/json"
	"net/http"
	
	"github.com/zlobste/mint-rest-api/internal/app/apiserver/helpers"
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

// UserSignUp handler for creating new users
func (s *server) SignUp() http.HandlerFunc {
	
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &models.User{
			Email: req.Email,
			Password:req.Password,
		}
		if err := s.store.User().Create(u); err != nil {
			helpers.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		u.Sanitize()
		helpers.Respond(w, r, http.StatusCreated, u)
	}
}
