package store

import (
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

type UserRepository interface {
	Create(model *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindById(id int64) (*models.User, error)
}

type OrderRepository interface {
	Create(model *models.Order) error
	FindById(id int64) (*models.Order, error)
	Cancel(id int64) error
}

type DishRepository interface {
	Create(model *models.Dish) error
	FindById(id int64) (*models.Dish, error)
	DeleteById(id int64) error
	GetAllDishes() ([]models.Dish, error)
}

type PaymentDetailsRepository interface {
	Create(model *models.PaymentDetails) error
	FindById(id int64) (*models.PaymentDetails, error)
	DeleteById(id int64) error
}
