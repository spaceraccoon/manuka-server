package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spaceraccoon/manuka-server/models"
)

// GetHoneypots gets all honeypots and returns as JSON
func GetHoneypots(c *gin.Context) {
	var Honeypots []models.Honeypot
	err := models.GetHoneypots(&Honeypots)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, Honeypots)
	return
}
