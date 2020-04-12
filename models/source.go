package models

import (
	"fmt"
	"time"

	"github.com/spaceraccoon/manuka-server/config"
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
	Name      string     `gorm:"not null" json:"name"`
	Type      uint       `json:"type"`
	APIKey    string     `json:"apiKey"`
	Honeypots []Honeypot `json:"honeypots"`
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
