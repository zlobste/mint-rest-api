package model

type PaymentDetails struct {
	Id              int64   `json:"id"`
	Bank            string  `json:"bank"`
	Account         string  `json:"account"`
	OrganizationId  string  `json:"organization_id"`
}