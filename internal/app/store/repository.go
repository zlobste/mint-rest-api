package store

import (
	"github.com/zlobste/mint-rest-api/internal/app/model"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindById(id int64) (*model.User, error)
}
