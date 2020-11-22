package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Dish struct {
	Id              int64   `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Cost            float64 `json:"cost"`
	MenuId          int64  `json:"menu_id"`
}

func (d *Dish) Validate() error {
	return validation.ValidateStruct(
		d,
		validation.Field(&d.Name, validation.Required, validation.Length(6, 100)),
		validation.Field(&d.Description, validation.Required, validation.Length(6, 2000)),
		validation.Field(&d.Cost, validation.Required),
		validation.Field(&d.Cost, validation.Required),
	)
}