package models

import (
	"errors"
	
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
	
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Password    string  `json:"password,omitempty"`
	Role        string  `json:"role"`
	Balance     string  `json:"balance"`
	Blocked     bool    `json:"blocked"`
}

func (u *User) EncryptPassword() error {
	if len(u.Password) == 0 {
		return errors.New("Wrong password length!")
	}
	enc, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.Password = string(enc)
	return nil
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Name, validation.Required, validation.Length(6, 100)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 100)),
		validation.Field(&u.Role, validation.Required),
	)
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(u.Password)) != nil
}
