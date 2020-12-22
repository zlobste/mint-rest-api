package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"

	"github.com/zlobste/mint-rest-api/internal/app/apiserver/server/helpers"
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

func (server *server) CreateOrder() http.HandlerFunc {
	type request struct {
		Cost     float64   `json:"cost"`
		DateTime time.Time `json:"datetime"`
		DishId   int64     `json:"dish_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tokenAuth, err := helpers.ExtractTokenMetadata(r)
		if err != nil {
			helpers.Error(w, r, http.StatusUnauthorized, err)
			return
		}

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		o := &models.Order{
			Cost:     req.Cost,
			DateTime: req.DateTime,
			DishId:   req.DishId,
			UserId:   tokenAuth.UserId,
		}
		if err := server.store.Order().Create(o); err != nil {
			helpers.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		helpers.Respond(w, r, http.StatusCreated, o)
	}
}

func (server *server) GetOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		order, err := server.store.Order().FindById(int64(id))
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, order)
	}
}

func (server *server) CancelOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		err := server.store.Order().CancelOrder(int64(id))
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, nil)
	}
}

func (server *server) SetStatusReady() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		err := server.store.Order().SetStatusReady(int64(id))
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

func (server *server) GetAllOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenAuth, err := helpers.ExtractTokenMetadata(r)
		if err != nil {
			helpers.Error(w, r, http.StatusUnauthorized, err)
			return
		}

		orders, err := server.store.Order().GetAllOrders(tokenAuth.UserId)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, orders)
	}
}
