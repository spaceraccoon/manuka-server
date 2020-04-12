package models

import (
	"fmt"
	"time"

	"github.com/spaceraccoon/manuka-server/config"
)

// Listener model
type Listener struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Name      string     `json:"name"`
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
