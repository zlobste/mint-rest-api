package model

import (
	"time"
)

type Order struct {
	Id              int64       `json:"id"`
	DishId          string      `json:"dish_id"`
	UserId          string      `json:"user_id"`
	Cost            float64     `json:"cost"`
	DateTime        time.Time   `json:"datetime"`
}