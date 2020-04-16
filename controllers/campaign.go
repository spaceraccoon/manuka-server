package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spaceraccoon/manuka-server/models"
)

// GetCampaigns gets all campaigns and returns as JSON
func GetCampaigns(c *gin.Context) {
	var Campaigns []models.Campaign
	err := models.GetCampaigns(&Campaigns)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, Campaigns)
	return
}

// CreateCampaign creates a campaign and returns as JSON
func CreateCampaign(c *gin.Context) {
	var campaign models.Campaign
	c.BindJSON(&campaign)
	if err := campaign.Validate(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if err := models.CreateCampaign(&campaign); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, campaign)
	return
}

// GetCampaign gets a campaign and returns as JSON
func GetCampaign(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var campaign models.Campaign
	err = models.GetCampaign(&campaign, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, campaign)
	return
}

// UpdateCampaign updates a campaign and returns as JSON
func UpdateCampaign(c *gin.Context) {
	var campaign models.Campaign
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = models.GetCampaign(&campaign, id)
	if err != nil {
		c.JSON(http.StatusNotFound, campaign)
		return
	}
	c.BindJSON(&campaign)
	if err := campaign.Validate(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	err = models.UpdateCampaign(&campaign, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, campaign)
	return
}

// DeleteCampaign deletes a campaign
func DeleteCampaign(c *gin.Context) {
	var campaign models.Campaign
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = models.DeleteCampaign(&campaign, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Status(http.StatusOK)
	return
}
