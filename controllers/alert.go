package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mickaelmagniez/elastic-alert/models"
	"net/http"
)

type AlertController struct{}

var alertModel = new(models.AlertModel)

func (ctrl AlertController) All(c *gin.Context) {
	alerts, err := alertModel.All()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"alerts": alerts})

	}

}

func (ctrl AlertController) Get(c *gin.Context) {
	id := c.Param("id")
	alert, err := alertModel.Get(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"alert": alert})
	}

}

func (ctrl AlertController) Create(c *gin.Context) {
	var alert models.Alert
	if err := c.ShouldBindJSON(&alert); err == nil {
		alert, err := alertModel.Create(alert)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"alert": alert})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (ctrl AlertController) Delete(c *gin.Context) {
	id := c.Param("id")
	id, err := alertModel.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"id": id})
	}

}
