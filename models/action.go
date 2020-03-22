package models

import "time"

// ActionType defines the different actions
type ActionType int

// Since sqlite doesn't support enums natively, we do it in golang
// However, we might want to use https://github.com/jinzhu/gorm/issues/1978 for postgres
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
