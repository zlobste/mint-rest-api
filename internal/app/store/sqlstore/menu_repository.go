package sqlstore

import (
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

type MenuRepository struct {
	store *Store
}

func (m *MenuRepository) Create(model *models.Menu) error {
	if err := model.Validate(); err != nil {
		return err
	}
	
	return m.store.db.QueryRow(
		"INSERT INTO menu (title, description, organization_id) VALUES ($1,$2, $3) RETURNING id",
		model.Title,
		model.Description,
		model.OrganizationId,
	).Scan(&model.Id)
}

func (m *MenuRepository) FindById(id int64) (*models.Menu, error) {
	model := &models.Menu{}
	if err := m.store.db.QueryRow("SELECT id, title, description, organization_id FROM menu WHERE i = $1", id).
		Scan(&model.Id, &model.Title, &model.Description, &model.OrganizationId); err != nil {
		return nil, err
	}
	return model, nil
}

func (m *MenuRepository)  DeleteById(id int64) error {
	_, err := m.store.db.Exec("DELETE FROM menu where id = $1", id)
	return err
}