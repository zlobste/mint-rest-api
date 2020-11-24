package models

import (
	"errors"
	
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
	
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	OWNER = iota
	ADMIN
	USER
)

type User struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Password    string  `json:"password,omitempty"`
	Role        int64   `json:"role"`
	Balance     string  `json:"balance"`
	Blocked     bool    `json:"blocked"`
}

func (user *User) EncryptPassword() error {
	if len(user.Password) == 0 {
		return errors.New("Wrong password length!")
	}
	enc, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	user.Password = string(enc)
	return nil
}

func (user *User) Validate() error {
	return validation.ValidateStruct(
		user,
		validation.Field(&user.Name, validation.Required, validation.Length(6, 100)),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 100)),
		validation.Field(&user.Role, validation.Required),
	)
}

func (user *User) Sanitize() {
	user.Password = ""
}

func (user *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)) != nil
}
