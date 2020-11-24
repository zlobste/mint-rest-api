package sqlstore

import (
	"errors"
	
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

type OrderRepository struct {
	store *Store
}

func (orderRepository *OrderRepository) Create(model *models.Order) error {
	if err := model.Validate(); err != nil {
		return err
	}
	
	return orderRepository.store.db.QueryRow(
		"INSERT INTO orders (cost, datetime, dish_id, user_id) VALUES ($1,$2, $3, $4) RETURNING id",
		model.Cost,
		model.DateTime,
		model.DishId,
		model.UserId,
	).Scan(&model.Id)
}

func (orderRepository *OrderRepository) FindById(id int64) (*models.Order, error) {
	model := &models.Order{}
	if err := orderRepository.store.db.QueryRow(
		"SELECT id, cost::decimal, datetime, status, dish_id, user_id FROM orders WHERE id=$1",
		id,
	).Scan(&model.Id, &model.Cost, &model.DateTime, &model.Status, &model.DishId, &model.UserId); err != nil {
		return nil, err
	}
	return model, nil
}


func (orderRepository *OrderRepository) CancelOrder(id int64) error {
	model := &models.Order{}
	if err := orderRepository.store.db.QueryRow("SELECT id, status FROM orders WHERE id=$1", id).
		Scan(&model.Id, &model.Status); err != nil {
		return err
	}
	
	if model.Status == models.PENDING {
		return errors.New("Order in progress!")
	} else if model.Status == models.REJECTED {
		return errors.New("Order is rejected!")
	}
	
	_, err := orderRepository.store.db.Exec("UPDATE orders SET status = $1 where id = $2",  models.REJECTED, id)
	return err
}

// For IOT
func (orderRepository *OrderRepository) SetStatusReady(id int64) error {
	model := &models.Order{}
	if err := orderRepository.store.db.QueryRow("SELECT id, status FROM orders WHERE id=$1", id).
		Scan(&model.Id, &model.Status); err != nil {
		return err
	}
	
	if model.Status == models.READY {
		return errors.New("Order is ready!")
	} else if model.Status == models.REJECTED {
		return errors.New("Order is rejected!")
	}
	
	_, err := orderRepository.store.db.Exec("UPDATE orders SET status = $1 where id = $2",  models.READY, id)
	return err
}

func (orderRepository *OrderRepository) GetOrderToExecute() (*models.Order, error) {
	rows, err := orderRepository.store.db.Query(
		"SELECT id, cost::decimal, datetime, status, dish_id, user_id FROM orders WHERE status = 0 ORDER BY datetime ASC",
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	orders := []*models.Order{}
	
	for rows.Next(){
		model := models.Order{}
		err := rows.Scan(&model.Id, &model.Cost, &model.DateTime, &model.Status, &model.DishId, &model.UserId)
		if err != nil{
			return  nil, err
		}
		orders = append(orders, &model)
	}
	
	return orders[0], nil
}

func (orderRepository *OrderRepository) CalculateSale(userId uint64, dishId uint64) (float64, error) {
	var average float64
	
	if err := orderRepository.store.db.QueryRow(
		"SELECT AVG(cost::decimal) FROM orders WHERE status = 2 AND user_id = $1",
		userId,
	).Scan(&average); err != nil {
		return 0, err
	}
	
	dish := &models.Dish{}
	if err := orderRepository.store.db.QueryRow(
		"SELECT cost::decimal, disabled FROM dishes WHERE id=$1",
		dishId,
	).Scan(&dish.Cost, &dish.Disabled); err != nil {
		return 0, err
	}
	
	if (dish.Disabled == true) {
		return 0, errors.New("Disabled dish!")
	}
	
	var sale float64 = 0;
	if (dish.Cost > average) {
		sale = dish.Cost * 0.1
	}
	
	return sale, nil
}