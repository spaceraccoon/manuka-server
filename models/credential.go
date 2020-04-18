package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spaceraccoon/manuka-server/config"
	"syreclabs.com/go/faker"
)

var (
	errCredentialUsernameRequired   = fmt.Errorf("Credential username required")
	errCredentialPasswordRequired   = fmt.Errorf("Credential password required")
	errCredentialHoneypotIDRequired = fmt.Errorf("Credential honeypot ID required")
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
					return errCredentialUsernameRequired
				}
			case "Password":
				switch validationErr.ActualTag() {
				case "required":
					return errCredentialPasswordRequired
				}
			case "HoneypotID":
				switch validationErr.ActualTag() {
				case "required":
					return errCredentialHoneypotIDRequired
				}
			default:
				return err
			}
		}
	}

	return nil
}

// BeforeSave hook validates credential
func (c *Credential) BeforeSave() (err error) {
	return c.Validate()
}

// CreateFakeCreds creates randomly-generated fake credentials and returns a string with the data
func CreateFakeCreds(h *Honeypot, source *Source) (credentials string, err error) {
	// Create fake credentials in the database and append to text
	n := rand.Intn(20) + 10
	var credential Credential
	credentials = ""
	for i := 0; i < n; i++ {
		credential = Credential{
			Username:   faker.Internet().FreeEmail(),
			Password:   faker.Internet().Password(8, 16),
			HoneypotID: h.ID,
		}
		err := CreateCredential(&credential)
		if err != nil {
			return "", err
		}
		credentials = credentials + credential.Username + ":" + credential.Password + "\n"
	}
	return credentials, nil
}

// CreateCredential creates a source in the database
func CreateCredential(credential *Credential) (err error) {
	if err = config.DB.Create(credential).Error; err != nil {
		return err
	}
	return nil
}
