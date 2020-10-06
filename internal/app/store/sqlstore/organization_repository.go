package sqlstore

import (
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

type OrganizationRepository struct {
	store *Store
}

func ( o *OrganizationRepository) Create(model *models.Organization) error {
	if err := model.Validate(); err != nil {
		return err
	}
	
	if err := model.EncryptPassword(); err != nil {
		return err
	}
	
	return o.store.db.QueryRow(
		"INSERT INTO organizations (name, email, password) VALUES ($1,$2, $3) RETURNING id",
		model.Name,
		model.Email,
		model.Password,
	).Scan(&model.Id)
}

func ( o *OrganizationRepository) FindByEmail(email string) (*models.Organization, error) {
	model := &models.Organization{}
	if err := o.store.db.QueryRow("SELECT id, name, email, password FROM organizations WHERE email=$1", email).
		Scan(&model.Id, &model.Name, &model.Email, &model.Password); err != nil {
		return nil, err
	}
	return model, nil
}

func ( o *OrganizationRepository) FindById(id int64) (*models.Organization, error) {
	model := &models.Organization{}
	if err := o.store.db.QueryRow("SELECT id, name, email, password FROM organizations WHERE id=$1", id).
		Scan(&model.Id, &model.Name, &model.Email, &model.Password); err != nil {
		return nil, err
	}
	return model, nil
}