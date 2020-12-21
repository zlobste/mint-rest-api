package sqlstore

import (
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

type UserRepository struct {
	store *Store
}

func (UserRepository *UserRepository) Create(model *models.User) error {
	if err := model.Validate(); err != nil {
		return err
	}

	if err := model.EncryptPassword(); err != nil {
		return err
	}

	return UserRepository.store.db.QueryRow(
		"INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id",
		model.Name,
		model.Email,
		model.Password,
		model.Role,
	).Scan(&model.Id)
}

func (UserRepository *UserRepository) FindByEmail(email string) (*models.User, error) {
	model := &models.User{}
	if err := UserRepository.store.db.QueryRow(
		"SELECT id, name, email, password, role, balance, blocked FROM users WHERE email=$1",
		email,
	).Scan(&model.Id, &model.Name, &model.Email, &model.Password, &model.Role, &model.Balance, &model.Blocked);
		err != nil {
		return nil, err
	}
	return model, nil
}

func (UserRepository *UserRepository) FindById(id int64) (*models.User, error) {
	model := &models.User{}
	if err := UserRepository.store.db.QueryRow(
		"SELECT id, name, email, password, role, balance, blocked FROM users WHERE id=$1",
		id,
	).Scan(&model.Id, &model.Name, &model.Email, &model.Password, &model.Role, &model.Balance, &model.Blocked);
		err != nil {
		return nil, err
	}
	return model, nil
}

func (UserRepository *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := UserRepository.store.db.Query("SELECT id, name, email, balance, role, blocked FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	users := []models.User{}

	for rows.Next() {
		model := models.User{}
		err := rows.Scan(&model.Id, &model.Name, &model.Email, &model.Balance, &model.Role, &model.Blocked)
		if err != nil {
			return nil, err
		}
		users = append(users, model)
	}

	return users, nil
}

func (UserRepository *UserRepository) BlockUser(id int64) error {
	_, err := UserRepository.store.db.Exec("UPDATE users SET blocked = TRUE WHERE id = $1", id)
	return err
}

func (UserRepository *UserRepository) UnblockUser(id int64) error {
	_, err := UserRepository.store.db.Exec("UPDATE users SET blocked = FALSE WHERE id = $1", id)
	return err
}
