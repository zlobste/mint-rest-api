package sqlstore

import (
	"database/sql"
	
	"github.com/zlobste/mint-rest-api/internal/app/store"
)

type Store struct {
	db                      *sql.DB
	userRepository          *UserRepository
	organizationRepository  *OrganizationRepository
	menuRepository          *MenuRepository
	orderRepository         *OrderRepository
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

func (s *Store) Organization() store.OrganizationRepository {
	if s.organizationRepository != nil {
		return s.organizationRepository
	}
	s.organizationRepository = &OrganizationRepository{
		store: s,
	}
	return s.organizationRepository
}

func (s *Store) Menu() store.MenuRepository {
	if s.menuRepository != nil {
		return s.menuRepository
	}
	s.menuRepository = &MenuRepository{
		store: s,
	}
	return s.menuRepository
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