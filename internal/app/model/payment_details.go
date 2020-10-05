package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type PaymentDetails struct {
	Id              int64   `json:"id"`
	Bank            string  `json:"bank"`
	Account         string  `json:"account"`
	OrganizationId  string  `json:"organization_id"`
}

func (p *PaymentDetails) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Bank, validation.Required),
		validation.Field(&p.Account, validation.Required),
		validation.Field(&p.OrganizationId, validation.Required),
	)
}