package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	errHoneypotNameRequired       = fmt.Errorf("Honeypot name required")
	errHoneypotListenerIDRequired = fmt.Errorf("Honeypot listener ID required")
	errHoneypotSourceIDRequired   = fmt.Errorf("Honeypot source ID required")
)

// Honeypot model
type Honeypot struct {
	ID          uint         `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DeletedAt   *time.Time   `json:"deletedAt"`
	Name        string       `json:"name" validate:"required"`
	CampaignID  uint         `json:"campaignId"`
	Credentials []Credential `json:"credentials"`
	ListenerID  uint         `json:"listenerId" validate:"required"`
	SourceID    uint         `json:"sourceId" validate:"required"`
}

// Validate validates struct fields
func (h *Honeypot) Validate() error {
	validate := validator.New()
	if err := validate.Struct(h); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, validationErr := range err.(validator.ValidationErrors) {
			switch validationErr.StructField() {
			case "Name":
				switch validationErr.ActualTag() {
				case "required":
					return errHoneypotNameRequired
				}
			case "ListenerID":
				switch validationErr.ActualTag() {
				case "required":
					return errHoneypotListenerIDRequired
				}
			case "SourceID":
				switch validationErr.ActualTag() {
				case "required":
					return errHoneypotSourceIDRequired
				}
			default:
				return err
			}
		}
	}

	if len(h.Credentials) > 0 {
		for idx := range h.Credentials {
			if err := h.Credentials[idx].Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
