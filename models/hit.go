package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/spaceraccoon/manuka-server/config"
)

var (
	ErrHitCredentialIDRequired = fmt.Errorf("hit credential id required")
	ErrHitHoneypotIDRequired   = fmt.Errorf("hit honeypot id required")
	ErrHitListenerIDRequired   = fmt.Errorf("hit listener id required")
	ErrHitSourceIDRequired     = fmt.Errorf("hit source id required")
	ErrHitIPAddressRequired    = fmt.Errorf("hit ip address required")
)

// Hit model
type Hit struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt"`
	CredentialID uint       `json:"credentialId" validate:"required"`
	HoneypotID   uint       `json:"honeypotId" validate:"required"`
	ListenerID   uint       `json:"listenerId" validate:"required"`
	SourceID     uint       `json:"sourceId" validate:"required"`
	IPAddress    string     `json:"ipAddress" validate:"required"`
}

// Validate validates struct fields
func (h *Hit) Validate() error {
	validate := validator.New()
	if err := validate.Struct(h); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, validationErr := range err.(validator.ValidationErrors) {
			switch validationErr.StructField() {
			case "CredentialID":
				switch validationErr.ActualTag() {
				case "required":
					return ErrHitCredentialIDRequired
				}
			case "HoneypotID":
				switch validationErr.ActualTag() {
				case "required":
					return ErrHitHoneypotIDRequired
				}
			case "ListenerID":
				switch validationErr.ActualTag() {
				case "required":
					return ErrHitListenerIDRequired
				}
			case "SourceID":
				switch validationErr.ActualTag() {
				case "required":
					return ErrHitSourceIDRequired
				}
			case "IPAddress":
				switch validationErr.ActualTag() {
				case "required":
					return ErrHitIPAddressRequired
				}
			default:
				return err
			}
		}
	}

	return nil
}

// GetHits gets all hits in database
func GetHits(hits *[]Hit) (err error) {
	if err = config.DB.Find(&hits).Error; err != nil {
		return err
	}
	return nil
}

// GetHit gets a hit in the database corresponding to id
func GetHit(hit *Hit, id int64) (err error) {
	if err := config.DB.First(&hit, id).Error; err != nil {
		return err
	}
	return nil
}

// DeleteHit deletes a hit in the database
func DeleteHit(hit *Hit, id int64) (err error) {
	config.DB.Where("id = ?", id).Delete(hit)
	return nil
}
