package models

import "github.com/go-ozzo/ozzo-validation"

// Artist represents an artist record.
type Artist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Validate validates the Artist fields.
func (a Artist) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required, validation.Length(0, 120)),
	)
}
