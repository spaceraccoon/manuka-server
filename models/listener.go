package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/spaceraccoon/manuka-server/config"
)

var (
	errListenerEmailRequired = fmt.Errorf("Listener email required")
	errListenerNameRequired  = fmt.Errorf("Listener name required")
	errListenerTypeRequired  = fmt.Errorf("Listenter type required")
	errListenerURLRequired   = fmt.Errorf("Listener URL required")
	errListenerTypeInvalid   = fmt.Errorf("Listener type invalid")
)

// ListenerType defines the different listener types
type ListenerType int

// Enumerate various listeners
const (
	LoginListener ListenerType = iota + 1
	SocialListener
)

// Listener model
type Listener struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Name      string     `json:"name" validate:"required"`
	Type      uint       `json:"type" validate:"required"`
	URL       *string    `json:"url"`
	Email     *string    `json:"email"`
}

// Validate validates struct fields
func (l *Listener) Validate() error {
	validate := validator.New()
	if err := validate.Struct(l); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, validationErr := range err.(validator.ValidationErrors) {
			switch validationErr.StructField() {
			case "Name":
				switch validationErr.ActualTag() {
				case "required":
					return errListenerNameRequired
				}
			default:
				return err
			}
		}
	}

	switch ListenerType(l.Type) {
	case LoginListener:
		if err := validate.Var(*l.URL, "required"); err != nil {
			return errListenerURLRequired
		}
	case SocialListener:
		if err := validate.Var(*l.Email, "required"); err != nil {
			return errListenerEmailRequired
		}
	default:
		return errListenerTypeInvalid
	}

	return nil
}

// GetListeners gets all listeners in database
func GetListeners(listeners *[]Listener) (err error) {
	if err = config.DB.Find(&listeners).Error; err != nil {
		return err
	}
	return nil
}

// CreateListener creates a listener in the database
func CreateListener(listener *Listener) (err error) {
	if err = config.DB.Create(listener).Error; err != nil {
		return err
	}
	return nil
}

// GetListener gets a listener in the database corresponding to id
func GetListener(listener *Listener, id int64) (err error) {
	if err := config.DB.First(&listener, id).Error; err != nil {
		return err
	}
	return nil
}

// UpdateListener updates a listener in the database
func UpdateListener(listener *Listener, id int64) (err error) {
	fmt.Println(listener)
	config.DB.Model(&listener).Update(listener)
	return nil
}

// DeleteListener deletes a listener in the database
func DeleteListener(listener *Listener, id int64) (err error) {
	config.DB.Where("id = ?", id).Delete(listener)
	return nil
}
