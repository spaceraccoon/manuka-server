package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/spaceraccoon/manuka-server/config"
)

var (
	ErrCampaignNameRequired = fmt.Errorf("campaign name is required")
)

// Campaign model
type Campaign struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Name      string     `json:"name" validate:"required"`
	Honeypots []Honeypot `json:"honeypots"`
}

// Validate validates struct fields
func (c *Campaign) Validate() error {
	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, validationErr := range err.(validator.ValidationErrors) {
			switch validationErr.StructField() {
			case "Name":
				switch validationErr.ActualTag() {
				case "required":
					return ErrCampaignNameRequired
				}
			default:
				return err
			}
		}
	}

	return nil
}

// GetCampaigns gets all campaigns in database
func GetCampaigns(campaigns *[]Campaign) (err error) {
	if err = config.DB.Preload("Honeypots").Find(&campaigns).Error; err != nil {
		return err
	}
	return nil
}

// CreateCampaign creates a campaign in the database
func CreateCampaign(campaign *Campaign) (err error) {
	if err = config.DB.Create(campaign).Error; err != nil {
		return err
	}
	return nil
}

// GetCampaign gets a campaign in the database corresponding to id
func GetCampaign(campaign *Campaign, id int64) (err error) {
	if err := config.DB.Preload("Honeypots").First(&campaign, id).Error; err != nil {
		return err
	}
	return nil
}

// UpdateCampaign updates a campaign in the database
func UpdateCampaign(campaign *Campaign, id int64) (err error) {
	fmt.Println(campaign)
	config.DB.Model(&campaign).Update(campaign)
	return nil
}

// DeleteCampaign deletes a campaign in the database
func DeleteCampaign(campaign *Campaign, id int64) (err error) {
	config.DB.Where("id = ?", id).Delete(campaign)
	return nil
}
