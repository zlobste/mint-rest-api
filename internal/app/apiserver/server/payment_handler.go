package server

import (
	"encoding/json"
	"net/http"
	
	"github.com/zlobste/mint-rest-api/internal/app/apiserver/server/helpers"
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

func (s *server) CreatePaymentDetails() http.HandlerFunc {
	type request struct {
		Bank            string  `json:"bank"`
		Account         string  `json:"account"`
		OrganizationId  string  `json:"organization_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		pd := &models.PaymentDetails{
			Bank: req.Bank,
			Account: req.Account,
			OrganizationId: req.OrganizationId,
		}
		if err := s.store.PaymentDetails().Create(pd); err != nil {
			helpers.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		helpers.Respond(w, r, http.StatusCreated, pd)
	}
}

func (s *server) DeletePaymentDetails() http.HandlerFunc {
	type request struct {
		id int64 `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		err := s.store.PaymentDetails().DeleteById(req.id)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) GetPaymentDetails() http.HandlerFunc {
	type request struct {
		id int64 `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}
		paymentDetails, err := s.store.PaymentDetails().FindById(req.id)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, err)
			return
		}

		helpers.Respond(w, r, http.StatusOK, paymentDetails)
	}
}
