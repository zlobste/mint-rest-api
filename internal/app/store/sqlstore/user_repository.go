package sqlstore

import (
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

type UserRepository struct {
	store *Store
}

func ( r *UserRepository) Create(model *models.User) error {
	if err := model.Validate(); err != nil {
		return err
	}
	
	if err := model.EncryptPassword(); err != nil {
		return err
	}
	
	return r.store.db.QueryRow(
		"INSERT INTO users (name, email, password) VALUES ($1,$2, $3) RETURNING id",
		model.Name,
		model.Email,
		model.Password,
	).Scan(&model.Id)
}

func ( r *UserRepository) FindByEmail(email string) (*models.User, error) {
	model := &models.User{}
	if err := r.store.db.QueryRow("SELECT id, name, email, password FROM users WHERE email=$1", email).
		Scan(&model.Id, &model.Name, &model.Email, &model.Password); err != nil {
		return nil, err
	}
	return model, nil
}

func ( r *UserRepository) FindById(id int64) (*models.User, error) {
	model := &models.User{}
	if err := r.store.db.QueryRow("SELECT id, name, email, password FROM users WHERE id=$1", id).
		Scan(&model.Id, &model.Name, &model.Email, &model.Password); err != nil {
		return nil, err
	}
	return model, nil
}