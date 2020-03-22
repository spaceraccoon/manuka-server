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
	} else {
		c.JSON(http.StatusOK, Campaigns)
	}
}

// CreateCampaign creates a campaign and returns as JSON
func CreateCampaign(c *gin.Context) {
	var campaign models.Campaign
	c.BindJSON(&campaign)
	err := models.CreateCampaign(&campaign)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, campaign)
	}
}

// GetCampaign gets a campaign and returns as JSON
func GetCampaign(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	var campaign models.Campaign
	err = models.GetCampaign(&campaign, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, campaign)
	}
}

// UpdateCampaign updates a campaign and returns as JSON
func UpdateCampaign(c *gin.Context) {
	var campaign models.Campaign
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	err = models.GetCampaign(&campaign, id)
	if err != nil {
		c.JSON(http.StatusNotFound, campaign)
	}
	c.BindJSON(&campaign)
	err = models.UpdateCampaign(&campaign, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, campaign)
	}
}

// DeleteCampaign deletes a campaign
func DeleteCampaign(c *gin.Context) {
	var campaign models.Campaign
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	err = models.DeleteCampaign(&campaign, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.Status(http.StatusOK)
	}
}
