package models

import (
	"time"
	
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	PENDING = iota
	READY
	REJECTED
)

type Order struct {
	Id              int64       `json:"id"`
	Cost            float64     `json:"cost"`
	DateTime        time.Time   `json:"datetime"`
	Status          int64       `json:"status"`
	DishId          string      `json:"dish_id"`
	UserId          string      `json:"user_id"`
}

func (order *Order) Validate() error {
	return validation.ValidateStruct(
		order,
		validation.Field(&order.Cost, validation.Required),
		validation.Field(&order.DateTime, validation.Required),
		validation.Field(&order.DishId, validation.Required),
		validation.Field(&order.UserId, validation.Required),
	)
}