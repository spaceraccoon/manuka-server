package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/spaceraccoon/manuka-server/config"
)

var (
	errHitCredentialIDRequired = fmt.Errorf("Hit credential ID required")
	errHitHoneypotIDRequired   = fmt.Errorf("Hit honeypot ID required")
	errHitIPAddressRequired    = fmt.Errorf("Hit IP address required")
	errHitListenerIDRequired   = fmt.Errorf("Hit listener ID required")
	errHitSourceIDRequired     = fmt.Errorf("Hit source ID required")
	errHitTypeRequired         = fmt.Errorf("Hit type required")
)

// HitType defines the different hits
type HitType int

// Enumerate various actions
const (
	FacebookRequest HitType = iota + 1
	LoginAttempt
	LinkedInRequest
	LinkedInMessage
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
	Type         HitType    `json:"type" validate:"required"`
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
					return errHitCredentialIDRequired
				}
			case "HoneypotID":
				switch validationErr.ActualTag() {
				case "required":
					return errHitHoneypotIDRequired
				}
			case "IPAddress":
				switch validationErr.ActualTag() {
				case "required":
					return errHitIPAddressRequired
				}
			case "ListenerID":
				switch validationErr.ActualTag() {
				case "required":
					return errHitListenerIDRequired
				}
			case "SourceID":
				switch validationErr.ActualTag() {
				case "required":
					return errHitSourceIDRequired
				}
			case "Type":
				switch validationErr.ActualTag() {
				case "required":
					return errHitTypeRequired
				}
			default:
				return err
			}
		}
	}

	return nil
}

// BeforeSave hook validates hit
func (h *Hit) BeforeSave() (err error) {
	return h.Validate()
}

// GetHits gets all hits in database
func GetHits(hits *[]Hit) (err error) {
	if err = config.DB.Find(&hits).Error; err != nil {
		return err
	}
	return nil
}

// CreateHit creates a hit in the database
func CreateHit(hit *Hit) (err error) {
	if err = config.DB.Create(hit).Error; err != nil {
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
