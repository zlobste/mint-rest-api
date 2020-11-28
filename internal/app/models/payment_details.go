package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type PaymentDetails struct {
	Id            int64  `json:"id"`
	Bank          string `json:"bank"`
	Account       string `json:"account"`
	InstitutionId string `json:"institution_id"`
}

func (paymentDetails *PaymentDetails) Validate() error {
	return validation.ValidateStruct(
		paymentDetails,
		validation.Field(&paymentDetails.Bank, validation.Required),
		validation.Field(&paymentDetails.Account, validation.Required),
		validation.Field(&paymentDetails.InstitutionId, validation.Required),
	)
}
