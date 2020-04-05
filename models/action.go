package models

import "time"

// ActionType defines the different actions
type ActionType int

// Enumerate various actions
const (
	Gist ActionType = iota
	Pastebin
	Linkedin
)

// Action model
type Action struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Name      string     `json:"name"`
	Type      uint       `json:"type"`
}
