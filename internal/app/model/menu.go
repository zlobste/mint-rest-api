package model

type Menu struct {
	Id              int64   `json:"id"`
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	OrganizationId  string  `json:"organization_id"`
}