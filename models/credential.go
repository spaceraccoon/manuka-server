package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	ErrCredentialUsernameRequired   = fmt.Errorf("credential username required")
	ErrCredentialPasswordRequired   = fmt.Errorf("credential password required")
	ErrCredentialHoneypotIDRequired = fmt.Errorf("credential honeypot id required")
)

// Credential model
type Credential struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
	Username   string     `json:"username" validate:"required"`
	Password   string     `json:"password" validate:"required"`
	HoneypotID uint       `json:"honeypotId" validate:"required"`
}

// Validate validates struct fields
func (c *Credential) Validate() error {
	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, validationErr := range err.(validator.ValidationErrors) {
			switch validationErr.StructField() {
			case "Username":
				switch validationErr.ActualTag() {
				case "required":
					return ErrCredentialUsernameRequired
				}
			case "Password":
				switch validationErr.ActualTag() {
				case "required":
					return ErrCredentialPasswordRequired
				}
			case "HoneypotID":
				switch validationErr.ActualTag() {
				case "required":
					return ErrCredentialHoneypotIDRequired
				}
			default:
				return err
			}
		}
	}

	return nil
}
