package models

import (
	"time"

	"github.com/spaceraccoon/manuka-server/config"
)

// Hit model
type Hit struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt"`
	CredentialID uint       `json:"credentialId"`
	HoneypotID   uint       `json:"honeypotId"`
	ListenerID   uint       `json:"listenerId"`
	SourceID     uint       `json:"sourceId"`
	IPAddress    string     `json:"ipAddress;not null"`
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
