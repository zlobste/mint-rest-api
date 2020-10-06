package server

import (
	"encoding/json"
	"errors"
	"net/http"
	
	"github.com/zlobste/mint-rest-api/internal/app/apiserver/helpers"
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not authenticated")
)

func (s *server) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenAuth, err := helpers.ExtractTokenMetadata(r)
		if err != nil {
			helpers.Error(w, r, http.StatusUnauthorized, err)
			return
		}
		user, err := s.store.User().FindById(tokenAuth.UserId)
		if err != nil {
			helpers.Error(w, r, http.StatusUnauthorized, err)
			return
		}
		user.Sanitize()
		helpers.Respond(w, r, http.StatusCreated, user)
	}
}

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