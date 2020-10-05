package sqlstore

import (
	"github.com/zlobste/mint-rest-api/internal/app/model"
)

type OrderRepository struct {
	store *Store
}

func (o *OrderRepository) Create(model *model.Order) error {
	if err := model.Validate(); err != nil {
		return err
	}
	
	return o.store.db.QueryRow(
		"INSERT INTO order (cost, datetime, dish_id, user_id) VALUES ($1,$2, $3, $4) RETURNING id",
		model.Cost,
		model.DateTime,
		model.DishId,
		model.UserId,
	).Scan(&model.Id)
}

func ( o *OrderRepository) FindById(id int64) (*model.Order, error) {
	model := &model.Order{}
	if err := o.store.db.QueryRow("SELECT id, title, description, organization_id FROM order WHERE id=$1", id).
		Scan(&model.Id, &model.Cost, &model.DateTime, &model.DishId, &model.UserId); err != nil {
		return nil, err
	}
	return model, nil
}
