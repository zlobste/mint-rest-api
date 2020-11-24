package sqlstore

import (
	"errors"
	
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

type OrderRepository struct {
	store *Store
}

func (o *OrderRepository) Create(model *models.Order) error {
	if err := model.Validate(); err != nil {
		return err
	}
	
	return o.store.db.QueryRow(
		"INSERT INTO orders (cost, datetime, dish_id, user_id) VALUES ($1,$2, $3, $4) RETURNING id",
		model.Cost,
		model.DateTime,
		model.DishId,
		model.UserId,
	).Scan(&model.Id)
}

func (o *OrderRepository) FindById(id int64) (*models.Order, error) {
	model := &models.Order{}
	if err := o.store.db.QueryRow("SELECT id, cost, datetime, status, dish_id, user_id FROM orders WHERE id=$1", id).
		Scan(&model.Id, &model.Cost, &model.DateTime, &model.Status, &model.DishId, &model.UserId); err != nil {
		return nil, err
	}
	return model, nil
}


func (o *OrderRepository) Cancel(id int64) error {
	model := &models.Order{}
	if err := o.store.db.QueryRow("SELECT id, status FROM orders WHERE id=$1", id).
		Scan(&model.Id, &model.Status); err != nil {
		return err
	}
	
	if model.Status == models.PENDING {
		return errors.New("Order in progress!")
	} else if model.Status == models.REJECTED {
		return errors.New("Order is rejected!")
	}
	
	_, err := o.store.db.Exec("UPDATE orders SET status = $1 where id = $2",  models.REJECTED, id)
	return err
}

// For IOT
func (o *OrderRepository) Ready(id int64) error {
	model := &models.Order{}
	if err := o.store.db.QueryRow("SELECT id, status FROM orders WHERE id=$1", id).
		Scan(&model.Id, &model.Status); err != nil {
		return err
	}
	
	if model.Status == models.READY {
		return errors.New("Order is ready!")
	} else if model.Status == models.REJECTED {
		return errors.New("Order is rejected!")
	}
	
	_, err := o.store.db.Exec("UPDATE orders SET status = $1 where id = $2",  models.READY, id)
	return err
}

