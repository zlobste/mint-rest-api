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

func (i *Institution) Validate() error {
	return validation.ValidateStruct(
		i,
		validation.Field(&i.Title, validation.Required),
		validation.Field(&i.Address, validation.Required),
	)
}