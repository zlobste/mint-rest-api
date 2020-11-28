package sqlstore

import (
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

type PaymentDetailsRepository struct {
	store *Store
}

func (paymentDetailsRepository *PaymentDetailsRepository) Create(model *models.PaymentDetails) error {
	if err := model.Validate(); err != nil {
		return err
	}
	
	return paymentDetailsRepository.store.db.QueryRow(
		"INSERT INTO payment_details (bank, account, institution_id) VALUES ($1,$2, $3) RETURNING id",
		model.Bank,
		model.Account,
		model.InstitutionId,
	).Scan(&model.Id)
}

func (paymentDetailsRepository *PaymentDetailsRepository) FindById(id int64) (*models.PaymentDetails, error) {
	model := &models.PaymentDetails{}
	if err := paymentDetailsRepository.store.db.QueryRow(
		"SELECT id, bank, account, institution_id FROM payment_details WHERE id=$1",
		id,
	).Scan(&model.Id, &model.Bank, &model.Account, &model.InstitutionId); err != nil {
		return nil, err
	}
	return model, nil
}

func (paymentDetailsRepository *PaymentDetailsRepository) DeleteById(id int64) error {
	_, err := paymentDetailsRepository.store.db.Exec("DELETE FROM payment_details WHERE id = $1", id)
	return err
}
