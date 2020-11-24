package sqlstore

import (
	"errors"
	
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

type UserRepository struct {
	store *Store
}

func (userRepository *UserRepository) Create(model *models.User) error {
	if err := model.Validate(); err != nil {
		return err
	}
	
	if err := model.EncryptPassword(); err != nil {
		return err
	}
	
	return userRepository.store.db.QueryRow(
		"INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id",
		model.Name,
		model.Email,
		model.Password,
		model.Role,
	).Scan(&model.Id)
}

func (userRepository *UserRepository) FindByEmail(email string) (*models.User, error) {
	model := &models.User{}
	if err := userRepository.store.db.QueryRow(
		"SELECT id, name, email, password, role, balance, blocked FROM users WHERE email=$1",
		email,
	).Scan(&model.Id, &model.Name, &model.Email, &model.Password, &model.Role, &model.Balance, &model.Blocked);
	err != nil {
		return nil, err
	}
	return model, nil
}

func (userRepository *UserRepository) FindById(id int64) (*models.User, error) {
	model := &models.User{}
	if err := userRepository.store.db.QueryRow(
		"SELECT id, name, email, password, role, balance, blocked FROM users WHERE id=$1",
		id,
	).Scan(&model.Id, &model.Name, &model.Email, &model.Password, &model.Role, &model.Balance, &model.Blocked);
	err != nil {
		return nil, err
	}
	return model, nil
}

func (userRepository *UserRepository) BlockUser(id int64) error {
	_, err := userRepository.store.db.Exec("UPDATE users SET blocked = true WHERE id = $1", id)
	return err
}

func (userRepository *UserRepository) UnblockUser(id int64) error {
	_, err := userRepository.store.db.Exec("UPDATE users SET blocked = false WHERE id = $1", id)
	return err
}