package store

import (
	"github.com/zlobste/mint-rest-api/internal/app/store/sqlstore"
)

type Store interface {
	User() *sqlstore.UserRepository
}