package sqlstore

import (
	"github.com/zlobste/mint-rest-api/internal/app/model"
)

type PaymentDetailsRepository struct {
	store *Store
}

func (p *PaymentDetailsRepository) Create(model *model.PaymentDetails) error {
	if err := model.Validate(); err != nil {
		return err
	}

	return p.store.db.QueryRow(
		"INSERT INTO payment_details (bank, account, organization_id) VALUES ($1,$2, $3) RETURNING id",
		model.Bank,
		model.Account,
		model.OrganizationId,
	).Scan(&model.Id)
}

func (p *PaymentDetailsRepository) FindById(id int64) (*model.PaymentDetails, error) {
	model := &model.PaymentDetails{}
	if err := p.store.db.QueryRow("SELECT id, bank, account, organization_id FROM payment_details WHERE id=$1", id).
		Scan(&model.Id, &model.Bank, &model.Account, &model.OrganizationId); err != nil {
		return nil, err
	}
	return model, nil
}

func (p *PaymentDetailsRepository) DeleteById(id int64) error {
	_, err := p.store.db.Exec("DELETE FROM payment_details where id = $1", id)
	return err
}
