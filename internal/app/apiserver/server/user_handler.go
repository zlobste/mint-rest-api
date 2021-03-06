package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/zlobste/mint-rest-api/internal/app/apiserver/server/helpers"
)

func (server *server) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenAuth, err := helpers.ExtractTokenMetadata(r)
		if err != nil {
			helpers.Error(w, r, http.StatusUnauthorized, err)
			return
		}
		user, err := server.store.User().FindById(tokenAuth.UserId)
		if err != nil {
			helpers.Error(w, r, http.StatusUnauthorized, err)
			return
		}
		user.Sanitize()
		helpers.Respond(w, r, http.StatusCreated, user)
	}
}

func (server *server) GetAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := server.store.User().GetAllUsers()
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, users)
	}
}

func (server *server) BlockUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		blocked, _ := strconv.ParseBool(params["blocked"])

		err := server.store.User().BlockUser(int64(id), blocked)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, id)
	}
}
