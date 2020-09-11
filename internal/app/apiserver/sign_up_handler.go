package apiserver

import (
	"encoding/json"
	"net/http"
	
	"github.com/zlobste/mint-rest-api/internal/app/model"
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
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			Email: req.Email,
			Password:req.Password,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}
