package model

type Dish struct {
	Id              int64   `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Cost            float64 `json:"cost"`
	MenuId          string  `json:"menu_id"`
}