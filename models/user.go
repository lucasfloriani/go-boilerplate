package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Nome    string `json:"nome"`
	Usuario string `json:"usuario"`
}

// Validate validates the Artist fields.
func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Nome, validation.Required, validation.Length(0, 120)),
		validation.Field(&u.Usuario, validation.Required, validation.Length(0, 120)),
	)
}
