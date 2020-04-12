package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spaceraccoon/manuka-server/models"
)

// GetSources gets all sources and returns as JSON
func GetSources(c *gin.Context) {
	var Sources []models.Source
	err := models.GetSources(&Sources)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, Sources)
	return
}

// CreateSource creates a source and returns as JSON
func CreateSource(c *gin.Context) {
	var source models.Source
	c.BindJSON(&source)
	err := models.CreateSource(&source)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusCreated, source)
	return
}

// GetSource gets a source and returns as JSON
func GetSource(c *gin.Context) {
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	var source models.Source
	err = models.GetSource(&source, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, source)
	return
}

// UpdateSource updates a source and returns as JSON
func UpdateSource(c *gin.Context) {
	var source models.Source
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = models.GetSource(&source, id)
	if err != nil {
		c.JSON(http.StatusNotFound, source)
		return
	}
	c.BindJSON(&source)
	err = models.UpdateSource(&source, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, source)
	return
}

// DeleteSource deletes a source
func DeleteSource(c *gin.Context) {
	var source models.Source
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = models.DeleteSource(&source, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Status(http.StatusOK)
	return
}
