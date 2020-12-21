package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/zlobste/mint-rest-api/internal/app/apiserver/server/helpers"
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

func (server *server) CreatePaymentDetails() http.HandlerFunc {
	type request struct {
		Bank          string `json:"bank"`
		Account       string `json:"account"`
		InstitutionId string `json:"institutionId_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		pd := &models.PaymentDetails{
			Bank:          req.Bank,
			Account:       req.Account,
			InstitutionId: req.InstitutionId,
		}
		if err := server.store.PaymentDetails().Create(pd); err != nil {
			helpers.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		helpers.Respond(w, r, http.StatusCreated, pd)
	}
}

func (server *server) DeletePaymentDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		err := server.store.PaymentDetails().DeleteById(int64(id))
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, nil)
	}
}

func (server *server) GetPaymentDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		paymentDetails, err := server.store.PaymentDetails().FindById(int64(id))
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, paymentDetails)
	}
}
