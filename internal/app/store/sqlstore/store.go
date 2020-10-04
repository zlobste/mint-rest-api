package sqlstore

import (
	"database/sql"
	
	"github.com/zlobste/mint-rest-api/internal/app/store"
)

type Store struct {
	db                      *sql.DB
	userRepository          *UserRepository
	organizationRepository  *OrganizationRepository
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