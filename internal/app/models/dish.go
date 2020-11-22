package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Dish struct {
	Id              int64   `json:"id"`
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	Cost            float64 `json:"cost"`
	Disabled        bool    `json:"disabled"`
}

func (d *Dish) Validate() error {
	return validation.ValidateStruct(
		d,
		validation.Field(&d.Title, validation.Required, validation.Length(6, 100)),
		validation.Field(&d.Description, validation.Required, validation.Length(6, 2000)),
		validation.Field(&d.Cost, validation.Required),
		validation.Field(&d.Cost, validation.Required),
	)
}