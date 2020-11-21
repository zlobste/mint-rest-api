package server

import (
	"encoding/json"
	"errors"
	"net/http"
	
	"github.com/zlobste/mint-rest-api/internal/app/apiserver/server/helpers"
	"github.com/zlobste/mint-rest-api/internal/app/models"
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
		Role string		`json:"role"`
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

		var id int64
		if req.Role == "user" {
			user, err := s.store.User().FindByEmail(req.Email)
			if err != nil || !user.ComparePassword(req.Password) {
				helpers.Error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
				return
			}
			id = user.Id
		} else {
			organization, err := s.store.User().FindByEmail(req.Email)
			if err != nil || !organization.ComparePassword(req.Password) {
				helpers.Error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
				return
			}
			id = organization.Id
		}

		token, err := helpers.CreateJWT(id)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, &JWT{Token: token})
	}
}

// UserSignUp handler for creating new users
func (s *server) SignUp() http.HandlerFunc {
	type request struct {
		Email       string  `json:"email"`
		Name        string  `json:"name"`
		Password    string  `json:"password"`
		Role        string  `json:"role"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		if req.Role == "user" {
			u := &models.User{
				Email: req.Email,
				Name: req.Name,
				Password:req.Password,
			}
			if err := s.store.User().Create(u); err != nil {
				helpers.Error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			u.Sanitize()
			helpers.Respond(w, r, http.StatusCreated, u)
		} else {
			o := &models.Organization{
				Email: req.Email,
				Password:req.Password,
			}
			if err := s.store.Organization().Create(o); err != nil {
				helpers.Error(w, r, http.StatusUnprocessableEntity, err)
				return
			}
			o.Sanitize()
			helpers.Respond(w, r, http.StatusCreated, o)
		}
	}
}
