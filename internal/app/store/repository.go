package store

import (
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

type UserRepository interface {
	Create(model *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindById(id int64) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}

type OrderRepository interface {
	Create(model *models.Order) error
	FindById(id int64) (*models.Order, error)
	CancelOrder(id int64) error
	SetStatusReady(id int64) error
	GetOrderToExecute() (*models.Order, error)
}

type DishRepository interface {
	Create(model *models.Dish) error
	FindById(id int64) (*models.Dish, error)
	DeleteById(id int64) error
	GetAllDishes() ([]models.Dish, error)
	CalculateSale(userId int64, dishId int64) (float64, error)
}

type PaymentDetailsRepository interface {
	Create(model *models.PaymentDetails) error
	FindById(id int64) (*models.PaymentDetails, error)
	DeleteById(id int64) error
}

type InstitutionRepository interface {
	Create(model *models.Institution) error
	FindByTitle(title string) ([]models.Institution, error)
	GetAllInstitutions() ([]models.Institution, error)
	DeleteById(id int64) error
}
