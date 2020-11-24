package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/zlobste/mint-rest-api/internal/app/store"
)

type Store struct {
	db                          *sql.DB
	userRepository              *UserRepository
	orderRepository             *OrderRepository
	dishRepository              *DishRepository
	paymentDetailsRepository    *PaymentDetailsRepository
	institutionRepository       *InstitutionRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (store *Store) User() store.UserRepository {
	if store.userRepository != nil {
		return store.userRepository
	}
	store.userRepository = &UserRepository{
		store: store,
	}
	return store.userRepository
}

func (store *Store) Order() store.OrderRepository {
	if store.orderRepository != nil {
		return store.orderRepository
	}
	store.orderRepository = &OrderRepository{
		store: store,
	}
	return store.orderRepository
}

func (store *Store) Dish() store.DishRepository {
	if store.dishRepository != nil {
		return store.dishRepository
	}
	store.dishRepository = &DishRepository{
		store: store,
	}
	return store.dishRepository
}

func (store *Store) Institution() store.InstitutionRepository {
	if store.institutionRepository != nil {
		return store.institutionRepository
	}
	store.institutionRepository = &InstitutionRepository{
		store: store,
	}
	return store.institutionRepository
}

func (store *Store) PaymentDetails() store.PaymentDetailsRepository {
	if store.paymentDetailsRepository != nil {
		return store.paymentDetailsRepository
	}
	store.paymentDetailsRepository = &PaymentDetailsRepository{
		store: store,
	}
	return store.paymentDetailsRepository
}