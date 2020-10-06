package sqlstore

import (
	"github.com/zlobste/mint-rest-api/internal/app/model"
)

type DishRepository struct {
	store *Store
}

func (d *DishRepository) Create(model *model.Dish) error {
	if err := model.Validate(); err != nil {
		return err
	}
	
	return d.store.db.QueryRow(
		"INSERT INTO dishes (name, description, cost, menu_id) VALUES ($1,$2, $3, $4) RETURNING id",
		model.Name,
		model.Description,
		model.Cost,
		model.MenuId,
	).Scan(&model.Id)
}

func (d *DishRepository) FindById(id int64) (*model.Dish, error) {
	model := &model.Dish{}
	if err := d.store.db.QueryRow("SELECT id, title, description, organization_id FROM dishes WHERE id=$1", id).
		Scan(&model.Id, &model.Name, &model.Description, &model.Cost, &model.MenuId); err != nil {
		return nil, err
	}
	return model, nil
}

func (d *DishRepository) DeleteById(id int64) error {
	_, err := d.store.db.Exec("DELETE FROM dishes where id = $1", id)
	return err
}