package server

import (
	"encoding/json"
	"net/http"
	"time"
	
	"github.com/zlobste/mint-rest-api/internal/app/apiserver/server/helpers"
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

func (server *server) CreateOrder() http.HandlerFunc {
	type request struct {
		Cost            float64     `json:"cost"`
		DateTime        time.Time   `json:"datetime"`
		DishId          string      `json:"dish_id"`
		UserId          string      `json:"user_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		o := &models.Order{
			Cost: req.Cost,
			DateTime: req.DateTime,
			DishId: req.DishId,
			UserId: req.UserId,
		}
		if err := server.store.Order().Create(o); err != nil {
			helpers.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		helpers.Respond(w, r, http.StatusCreated, o)
	}
}

func (server *server) GetOrder() http.HandlerFunc {
	type request struct {
		id int64 `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		order, err := server.store.Order().FindById(req.id)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, order)
	}
}

func (server *server) CancelOrder() http.HandlerFunc {
	type request struct {
		id int64 `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		err := server.store.Order().CancelOrder(req.id)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		
		helpers.Respond(w, r, http.StatusOK, nil)
	}
}

func (server *server) SetStatusReady() http.HandlerFunc {
	type request struct {
		id int64 `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		err := server.store.Order().SetStatusReady(req.id)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		
		helpers.Respond(w, r, http.StatusOK, nil)
	}
}

func (server *server) GetOrderToExecute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		order, err := server.store.Order().GetOrderToExecute()
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		
		helpers.Respond(w, r, http.StatusOK, order)
	}
}