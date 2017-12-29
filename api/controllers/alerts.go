package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mickaelmagniez/elastic-alert/models"
	"net/http"
	"github.com/mickaelmagniez/elastic-alert/store"
)

type AlertsController struct{}

func (AlertsController) All(c *gin.Context) {
	alerts, err := store.AllAlerts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, alerts)
	}
}

func (AlertsController) Get(c *gin.Context) {
	id := c.Param("id")
	alert, err := store.GetAlert(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, alert)
	}
}

func (AlertsController) Create(c *gin.Context) {
	var alert models.Alert
	if err := c.ShouldBindJSON(&alert); err == nil {
		alert, err := store.CreateAlert(alert)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		} else {
			c.JSON(http.StatusOK, alert)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (AlertsController) Update(c *gin.Context) {
	var alert models.Alert
	if err := c.ShouldBindJSON(&alert); err == nil {
		alert, err := store.UpdateAlert(alert)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		} else {
			c.JSON(http.StatusOK, alert)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (AlertsController) Delete(c *gin.Context) {
	id := c.Param("id")
	id, err := store.DeleteAlert(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
}
