package models

import "time"

// Credential model
type Credential struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	HoneypotID uint       `json:"honeypotId"`
}
