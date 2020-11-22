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
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

func (s *Store) Order() store.OrderRepository {
	if s.orderRepository != nil {
		return s.orderRepository
	}
	s.orderRepository = &OrderRepository{
		store: s,
	}
	return s.orderRepository
}

func (s *Store) Dish() store.DishRepository {
	if s.dishRepository != nil {
		return s.dishRepository
	}
	s.dishRepository = &DishRepository{
		store: s,
	}
	return s.dishRepository
}

func (s *Store) PaymentDetails() store.PaymentDetailsRepository {
	if s.paymentDetailsRepository != nil {
		return s.paymentDetailsRepository
	}
	s.paymentDetailsRepository = &PaymentDetailsRepository{
		store: s,
	}
	return s.paymentDetailsRepository
}