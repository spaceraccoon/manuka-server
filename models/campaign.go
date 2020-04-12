package models

import (
	"fmt"
	"time"

	"github.com/spaceraccoon/manuka-server/config"
)

// Campaign model
type Campaign struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Name      string     `json:"name;not null"`
	Honeypots []Honeypot `json:"honeypots"`
}

// GetCampaigns gets all campaigns in database
func GetCampaigns(campaigns *[]Campaign) (err error) {
	if err = config.DB.Find(&campaigns).Error; err != nil {
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
	if err := config.DB.First(&campaign, id).Error; err != nil {
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
