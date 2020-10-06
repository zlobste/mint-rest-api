package store

import (
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

type UserRepository interface {
	Create(model *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindById(id int64) (*models.User, error)
}

type OrganizationRepository interface {
	Create(model *models.Organization) error
	FindByEmail(email string) (*models.Organization, error)
	FindById(id int64) (*models.Organization, error)
}

type MenuRepository interface {
	Create(model *models.Menu) error
	FindById(id int64) (*models.Menu, error)
	DeleteById(id int64) error
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
}

type PaymentDetailsRepository interface {
	Create(model *models.PaymentDetails) error
	FindById(id int64) (*models.PaymentDetails, error)
	DeleteById(id int64) error
}
