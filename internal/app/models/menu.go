package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Menu struct {
	Id              int64   `json:"id"`
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	OrganizationId  string  `json:"organization_id"`
}

func (m *Menu) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Title, validation.Required, validation.Length(6, 100)),
		validation.Field(&m.Description, validation.Required, validation.Length(6, 2000)),
	)
}