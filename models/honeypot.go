package models

import (
	"fmt"
	"net/url"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/spaceraccoon/manuka-server/config"
	"github.com/spaceraccoon/manuka-server/utils"
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

// BeforeSave hook validates honeypot
func (h *Honeypot) BeforeSave() (err error) {
	return h.Validate()
}

// GetHoneypots gets all hits in database
func GetHoneypots(honeypots *[]Honeypot) (err error) {
	if err = config.DB.Find(&honeypots).Error; err != nil {
		return err
	}
	return nil
}

// AfterCreate creates and save the fake Pastebin credentials if the Honeypot has a Pastebin Source
func (h *Honeypot) AfterCreate(scope *gorm.Scope) (err error) {
	var source Source
	config.DB.Model(h).Related(&source)
	if SourceType(source.Type) == PastebinSource {
		credentials, err := CreateFakeCreds(h, &source)
		if err != nil {
			return err
		}

		var listener Listener
		config.DB.Model(h).Related(&listener)

		u, err := url.Parse(*listener.URL)
		if err != nil {
			return err
		}

		paste := &utils.Paste{
			Text:   u.Host + "\n" + credentials, // only include URL host to bypass spam filter
			Name:   "Login Credentials",
			APIKey: *source.APIKey,
		}
		pastebinURL, err := utils.CreatePaste(paste)

		// append new pastebinURL to source
		config.DB.Model(&source).Update("PastebinURLs", append(source.PastebinURLs, pastebinURL))
		if err != nil {
			return err
		}
	}
	return
}

// BeforeUpdate creates and save the fake Pastebin credentials if the Honeypot has a Pastebin Source
func (h *Honeypot) BeforeUpdate() (err error) {
	var oldHoneypot Honeypot
	if err := config.DB.First(&oldHoneypot, h.ID).Error; err != nil {
		return err
	}
	if oldHoneypot.SourceID != h.SourceID {
		var source Source
		config.DB.Model(h).Related(&source)
		if SourceType(source.Type) == PastebinSource {
			credentials, err := CreateFakeCreds(h, &source)
			if err != nil {
				return err
			}

			var listener Listener
			config.DB.Model(h).Related(&listener)

			u, err := url.Parse(*listener.URL)
			if err != nil {
				return err
			}

			paste := &utils.Paste{
				Text:   u.Host + "\n" + credentials, // only include URL host to bypass spam filter
				Name:   "Login Credentials",
				APIKey: *source.APIKey,
			}
			pastebinURL, err := utils.CreatePaste(paste)

			// append new pastebinURL to source
			config.DB.Model(&source).Update("PastebinURLs", append(source.PastebinURLs, pastebinURL))
			if err != nil {
				return err
			}
		}
	}
	return
}
