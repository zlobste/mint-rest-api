package sqlstore

import (
	"github.com/zlobste/mint-rest-api/internal/app/model"
)

type MenuRepository struct {
	store *Store
}

func (m *MenuRepository) Create(model *model.Menu) error {
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

func ( m *MenuRepository) FindById(id int64) (*model.Menu, error) {
	model := &model.Menu{}
	if err := m.store.db.QueryRow("SELECT id, title, description, organization_id FROM menu WHERE id=$1", id).
		Scan(&model.Id, &model.Title, &model.Description, &model.OrganizationId); err != nil {
		return nil, err
	}
	return model, nil
}
