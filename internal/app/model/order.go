package model

import (
	"time"
	
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Order struct {
	Id              int64       `json:"id"`
	Cost            float64     `json:"cost"`
	DateTime        time.Time   `json:"datetime"`
	DishId          string      `json:"dish_id"`
	UserId          string      `json:"user_id"`
}

func (o *Order) Validate() error {
	return validation.ValidateStruct(
		o,
		validation.Field(&o.Cost, validation.Required),
		validation.Field(&o.DateTime, validation.Required),
	)
}