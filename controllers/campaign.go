package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spaceraccoon/manuka-server/models"
)

// GetCampaigns gets all campaigns and returns as JSON
func GetCampaigns(c *gin.Context) {
	var campaigns []models.Campaign
	err := models.GetCampaigns(&campaigns)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, campaigns)
	return
}

// CreateCampaign creates a campaign and returns as JSON
func CreateCampaign(c *gin.Context) {
	var campaign models.Campaign
	err := c.BindJSON(&campaign)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := models.CreateCampaign(&campaign); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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
	err = c.BindJSON(&campaign)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = models.UpdateCampaign(&campaign, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
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
