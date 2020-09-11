package sqlstore

import (
	"github.com/zlobste/mint-rest-api/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func ( r *UserRepository) Create(model *model.User) error {
	if err := model.Validate(); err != nil {
		return err
	}
	
	if err := model.EncryptPassword(); err != nil {
		return err
	}
	
	return r.store.db.QueryRow(
		"INSERT INTO users (email, password) VALUES ($1,$2) RETURNING id",
		model.Email,
		model.Password,
	).Scan(&model.Id)
}

func ( r *UserRepository) FindByEmail(email string) (*model.User, error) {
	model := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, email, password FROM users WHERE email=$1", email).Scan(&model.Id, &model.Email, &model.Password); err != nil {
		return nil, err
	}
	return model, nil
}
