package store

import (
	"github.com/zlobste/mint-rest-api/internal/app/model"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail (emain string) (*model.User, error)
}
