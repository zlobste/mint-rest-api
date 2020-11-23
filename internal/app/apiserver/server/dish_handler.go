package server

import (
	"encoding/json"
	"net/http"
	
	"github.com/zlobste/mint-rest-api/internal/app/apiserver/server/helpers"
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

func (s *server) CreateDish() http.HandlerFunc {
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
		if err := s.store.Dish().Create(d); err != nil {
			helpers.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		helpers.Respond(w, r, http.StatusCreated, d)
	}
}

func (s *server) DeleteDish() http.HandlerFunc {

	type request struct {
		id int64 `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		err := s.store.Dish().DeleteById(req.id)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) GetDish() http.HandlerFunc {
	type request struct {
		id int64 `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		dish, err := s.store.Dish().FindById(req.id)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, dish)
	}
}

func (s *server) GetAllDishes() http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {
		dishes, err := s.store.Dish().GetAllDishes()
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		
		helpers.Respond(w, r, http.StatusOK, dishes)
	}
}