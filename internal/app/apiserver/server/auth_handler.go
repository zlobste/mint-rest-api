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
func (server *server) SignIn() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     int64  `json:"role"`
	}

	type JWT struct {
		Token string `json:"token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		user, err := server.store.User().FindByEmail(req.Email)
		if err != nil || !user.ComparePassword(req.Password) {
			helpers.Error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		if user.Blocked {
			helpers.Error(w, r, http.StatusBadRequest, err)
		}

		token, err := helpers.CreateJWT(user.Id)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, &JWT{Token: token})
	}
}

// UserSignUp handler for creating new users
func (server *server) SignUp() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
		Role     int64  `json:"role"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &models.User{
			Email:    req.Email,
			Name:     req.Name,
			Password: req.Password,
			Role:     req.Role,
		}
		if err := server.store.User().Create(u); err != nil {
			helpers.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		u.Sanitize()
		helpers.Respond(w, r, http.StatusCreated, u)
	}
}
