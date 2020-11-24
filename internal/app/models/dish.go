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

func (dish *Dish) Validate() error {
	return validation.ValidateStruct(
		dish,
		validation.Field(&dish.Title, validation.Required, validation.Length(6, 100)),
		validation.Field(&dish.Description, validation.Required, validation.Length(6, 2000)),
		validation.Field(&dish.Cost, validation.Required),
		validation.Field(&dish.Cost, validation.Required),
	)
}