package models

import "time"

// Honeypot model
type Honeypot struct {
	ID          uint         `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DeletedAt   *time.Time   `json:"deletedAt"`
	Name        string       `json:"name"`
	CampaignID  uint         `json:"campaignId"`
	Credentials []Credential `json:"credentials"`
	ListenerID  uint         `json:"listenerId"`
	SourceID    uint         `json:"sourceId"`
}
