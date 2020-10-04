package model

import (
	"errors"
	
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
	
	validation  "github.com/go-ozzo/ozzo-validation/v4"
)

type Organization struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Password    string  `json:"password,omitempty"`
}

func (o *Organization) EncryptPassword() error {
	if len(o.Password) == 0 {
		return errors.New("Wrong password length!")
	}
	enc, err := bcrypt.GenerateFromPassword([]byte(o.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	o.Password = string(enc)
	return nil
}

func (o *Organization) Validate() error {
	return validation.ValidateStruct(
		o,
		validation.Field(&o.Name, validation.Required, validation.Length(6, 100)),
		validation.Field(&o.Email, validation.Required, is.Email),
		validation.Field(&o.Password, validation.Required, validation.Length(6, 100)),
	)
}

func (o *Organization) Sanitize() {
	o.Password = ""
}

func (o *Organization) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(o.Password)) != nil
}