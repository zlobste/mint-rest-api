package store

import (
	"github.com/zlobste/mint-rest-api/internal/app/model"
	
	"errors"
	
	crypt "golang.org/x/crypto"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	store *Store
}

func ( r *UserRepository) Create(model *model.User) (*model.User, error) {
	
	encyptedPassword, err := EncryptPassword(model.Password)
	if err != nil {
		return  model, err
	}
	
	if err := r.store.db.QueryRow("INSERT INTO users (email, password) VALUES ($1,$2) RETURNING id", model.Email, encyptedPassword).Scan(&model.Id); err != nil {
		return nil, err
	}
	return model, nil
}

func ( r *UserRepository) FindByEmail(email string) (*model.User, error) {
	model := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, email, password FROM users WHERE emasil=$1", email).Scan(&model.Id, &model.Email, &model.Password); err != nil {
		return nil, err
	}
	return model, nil
}

func EncryptPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("Wrong password length!")
	}
	enc, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(enc), nil
}