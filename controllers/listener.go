package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spaceraccoon/manuka-server/models"
)

// GetListeners gets all listeners and returns as JSON
func GetListeners(c *gin.Context) {
	var Listeners []models.Listener
	err := models.GetListeners(&Listeners)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, Listeners)
	return
}

// CreateListener creates a listener and returns as JSON
func CreateListener(c *gin.Context) {
	var listener models.Listener
	c.BindJSON(&listener)
	err := models.CreateListener(&listener)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusCreated, listener)
	return
}

// GetListener gets a listener and returns as JSON
func GetListener(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var listener models.Listener
	err = models.GetListener(&listener, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, listener)
	return
}

// UpdateListener updates a listener and returns as JSON
func UpdateListener(c *gin.Context) {
	var listener models.Listener
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = models.GetListener(&listener, id)
	if err != nil {
		c.JSON(http.StatusNotFound, listener)
		return
	}
	c.BindJSON(&listener)
	err = models.UpdateListener(&listener, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, listener)
	return
}

// DeleteListener deletes a listener
func DeleteListener(c *gin.Context) {
	var listener models.Listener
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = models.DeleteListener(&listener, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Status(http.StatusOK)
	return
}
