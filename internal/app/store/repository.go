package store

import (
	"github.com/zlobste/mint-rest-api/internal/app/model"
)

type UserRepository interface {
	Create(model *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindById(id int64) (*model.User, error)
}

type OrganizationRepository interface {
	Create(model *model.Organization) error
	FindByEmail(email string) (*model.Organization, error)
	FindById(id int64) (*model.Organization, error)
}

type MenuRepository interface {
	Create(model *model.Menu) error
	FindById(id int64) (*model.Menu, error)
}

type OrderRepository interface {
	Create(model *model.Order) error
	FindById(id int64) (*model.Order, error)
}

type DishRepository interface {
	Create(model *model.Dish) error
	FindById(id int64) (*model.Dish, error)
}