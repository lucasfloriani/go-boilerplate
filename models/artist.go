package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

// Artist represents an artist record.
type Artist struct {
	gorm.Model
	Nome string `json:"nome"`
}

// Validate validates the Artist fields,
// this is executed before Create and Update actions
// from gorm, like a Observer Design Pattern.
func (a Artist) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Nome, validation.Required, validation.Length(0, 120)),
	)
}
