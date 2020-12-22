package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/zlobste/mint-rest-api/internal/app/apiserver/server/helpers"
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

func (server *server) CreateInstitution() http.HandlerFunc {
	type request struct {
		Title   string `json:"title"`
		Address string `json:"address"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		i := &models.Institution{
			Title:   req.Title,
			Address: req.Address,
		}
		if err := server.store.Institution().Create(i); err != nil {
			helpers.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		helpers.Respond(w, r, http.StatusCreated, i)
	}
}

func (server *server) DeleteInstitution() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		err := server.store.Institution().DeleteById(int64(id))
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, id)
	}
}

func (server *server) FindInstitutionsByTitle() http.HandlerFunc {
	type request struct {
		Title string `json:"title"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		dish, err := server.store.Institution().FindByTitle(req.Title)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, dish)
	}
}

func (server *server) GetAllInstitutions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		institutions, err := server.store.Institution().GetAllInstitutions()
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, institutions)
	}
}
