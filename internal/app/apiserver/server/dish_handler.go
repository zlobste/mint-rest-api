package server

import (
	"encoding/json"
	"net/http"
	
	"github.com/zlobste/mint-rest-api/internal/app/apiserver/server/helpers"
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

func (server *server) CreateDish() http.HandlerFunc {
	type request struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Cost        float64 `json:"cost"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		d := &models.Dish{
			Title: req.Title,
			Description: req.Description,
			Cost: req.Cost,
		}
		if err := server.store.Dish().Create(d); err != nil {
			helpers.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		helpers.Respond(w, r, http.StatusCreated, d)
	}
}

func (server *server) DeleteDish() http.HandlerFunc {

	type request struct {
		id int64 `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		err := server.store.Dish().DeleteById(req.id)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, nil)
	}
}

func (server *server) GetDish() http.HandlerFunc {
	type request struct {
		id int64 `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		dish, err := server.store.Dish().FindById(req.id)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, dish)
	}
}

func (server *server) GetAllDishes() http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {
		dishes, err := server.store.Dish().GetAllDishes()
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		
		helpers.Respond(w, r, http.StatusOK, dishes)
	}
}

func (server *server) CalculateSale() http.HandlerFunc {
	type request struct {
		dishId int64 `json:"dishId"`
	}
	
	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		
		tokenAuth, err := helpers.ExtractTokenMetadata(r)
		if err != nil {
			helpers.Error(w, r, http.StatusUnauthorized, err)
			return
		}
		
		sale, err := server.store.Dish().CalculateSale(tokenAuth.UserId, req.dishId)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		
		helpers.Respond(w, r, http.StatusOK, sale)
	}
}