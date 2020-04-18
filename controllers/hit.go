package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spaceraccoon/manuka-server/config"
	"github.com/spaceraccoon/manuka-server/models"
)

// ListenerHit struct describes the hit requests sent by listeners
type ListenerHit struct {
	ListenerID   uint                `json:"listenerId"`
	ListenerType models.ListenerType `json:"listenerType"`
	IPAddress    string              `json:"ipAddress"`
	Username     string              `json:"username"`
	Password     string              `json:"password"`
	Email        string              `json:"email"`
	HitType      models.HitType      `json:"hitType"`
}

// GetHits gets all hits and returns as JSON
func GetHits(c *gin.Context) {
	var Hits []models.Hit
	err := models.GetHits(&Hits)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, Hits)
	return
}

// CreateHit creates a hit and returns as JSON
func CreateHit(c *gin.Context) {
	var listenerHit ListenerHit
	err := c.BindJSON(&listenerHit)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var hit models.Hit
	switch listenerHit.ListenerType {
	case models.LoginListener:
		var credential models.Credential
		if err := config.DB.Where("username = ? AND password = ?", listenerHit.Username, listenerHit.Password).First(&credential).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var honeypot models.Honeypot
		config.DB.Model(credential).Related(&honeypot)
		hit = models.Hit{
			CampaignID:   honeypot.CampaignID,
			CredentialID: credential.ID,
			HoneypotID:   credential.HoneypotID,
			ListenerID:   listenerHit.ListenerID,
			SourceID:     honeypot.SourceID,
			IPAddress:    listenerHit.IPAddress,
			Type:         listenerHit.HitType,
		}
	case models.SocialListener:
		// For now, there is no difference in proccessing of Social hits and Login hits
		var credential models.Credential
		if err := config.DB.Where("username = ?", listenerHit.Email).First(&credential).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var honeypot models.Honeypot
		config.DB.Model(credential).Related(&honeypot)
		hit = models.Hit{
			CampaignID:   honeypot.CampaignID,
			CredentialID: credential.ID,
			HoneypotID:   credential.HoneypotID,
			ListenerID:   listenerHit.ListenerID,
			SourceID:     honeypot.SourceID,
			IPAddress:    listenerHit.IPAddress,
			Type:         listenerHit.HitType,
		}
	default:
		log.Fatal("Environment variable LISTENER_TYPE must be one of login, social")
	}
	if err := models.CreateHit(&hit); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, hit)
	return
}

// GetHit gets a hit and returns as JSON
func GetHit(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var hit models.Hit
	err = models.GetHit(&hit, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, hit)
	return
}

// DeleteHit deletes a Hit
func DeleteHit(c *gin.Context) {
	var Hit models.Hit
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = models.DeleteHit(&Hit, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Status(http.StatusOK)
	return
}
