package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/spaceraccoon/manuka-server/config"
)

var (
	ErrSourceNameRequired   = fmt.Errorf("source name required")
	ErrSourceAPIKeyRequired = fmt.Errorf("source api key required")
)

// SourceType defines the different fake OSINT sources
type SourceType int

// Enumerate various actions
const (
	Pastebin SourceType = iota
	Facebook
)

// Source model
type Source struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Name      string     `gorm:"not null" json:"name" validate:"required"`
	Type      uint       `json:"type" validate:"required"`
	APIKey    string     `json:"apiKey"`
	Honeypots []Honeypot `json:"honeypots"`
}

// Validate validates struct fields
func (s *Source) Validate() error {
	validate := validator.New()
	if err := validate.Struct(s); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, validationErr := range err.(validator.ValidationErrors) {
			switch validationErr.StructField() {
			case "Name":
				switch validationErr.ActualTag() {
				case "required":
					return ErrSourceNameRequired
				}
			default:
				return err
			}
		}
	}

	switch SourceType(s.Type) {
	case Pastebin:
		if err := validate.Var(s.APIKey, "required"); err != nil {
			return ErrSourceAPIKeyRequired
		}
	case Facebook:
	}

	return nil
}

// GetSources gets all sources in database
func GetSources(sources *[]Source) (err error) {
	if err = config.DB.Find(&sources).Error; err != nil {
		return err
	}
	return nil
}

// CreateSource creates a source in the database
func CreateSource(source *Source) (err error) {
	if err = config.DB.Create(source).Error; err != nil {
		return err
	}
	return nil
}

// GetSource gets a source in the database corresponding to id
func GetSource(source *Source, id int64) (err error) {
	if err := config.DB.First(&source, id).Error; err != nil {
		return err
	}
	return nil
}

// UpdateSource updates a source in the database
func UpdateSource(source *Source, id int64) (err error) {
	fmt.Println(source)
	config.DB.Model(&source).Update(source)
	return nil
}

// DeleteSource deletes a source in the database
func DeleteSource(source *Source, id int64) (err error) {
	config.DB.Where("id = ?", id).Delete(source)
	return nil
}
