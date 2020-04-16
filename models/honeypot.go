package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	ErrHoneypotNameRequired       = fmt.Errorf("honeypot name required")
	ErrHoneypotCampaignIDRequired = fmt.Errorf("honeypot campaign id required")
	ErrHoneypotListenerIDRequired = fmt.Errorf("honeypot listener id required")
	ErrHoneypotSourceIDRequired   = fmt.Errorf("honeypot source id required")
)

// Honeypot model
type Honeypot struct {
	ID          uint         `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DeletedAt   *time.Time   `json:"deletedAt"`
	Name        string       `json:"name" validate:"required"`
	CampaignID  uint         `json:"campaignId" validate:"required"`
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
					return ErrHoneypotNameRequired
				}
			case "CampaignID":
				switch validationErr.ActualTag() {
				case "required":
					return ErrHoneypotCampaignIDRequired
				}
			case "ListenerID":
				switch validationErr.ActualTag() {
				case "required":
					return ErrHoneypotListenerIDRequired
				}
			case "SourceID":
				switch validationErr.ActualTag() {
				case "required":
					return ErrHoneypotSourceIDRequired
				}
			default:
				return err
			}
		}
	}

	return nil
}
