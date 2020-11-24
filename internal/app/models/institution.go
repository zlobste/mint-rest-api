package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Institution struct {
	Id          int64   `json:"id"`
	Title       string  `json:"title"`
	Address     string  `json:"address"`
	Disabled    bool    `json:"disabled"`
}

func (institution *Institution) Validate() error {
	return validation.ValidateStruct(
		institution,
		validation.Field(&institution.Title, validation.Required),
		validation.Field(&institution.Address, validation.Required),
	)
}