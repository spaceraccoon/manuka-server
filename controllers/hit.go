package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spaceraccoon/manuka-server/models"
)

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
