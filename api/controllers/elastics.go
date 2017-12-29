package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/mickaelmagniez/elastic-alert/store"
)

type ElasticsController struct{}

func (ElasticsController) GetServers(c *gin.Context) {
	servers, err := store.GetElasticServers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, servers)
	}
}

func (ElasticsController) GetIndices(c *gin.Context) {
	indices, err := store.GetElasticIndicesOfServer(c.Query("url"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, indices)
	}
}

func (ElasticsController) GetTypes(c *gin.Context) {
	types, err := store.GetElasticTypesOfIndex(c.Query("url"), c.Query("index"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, types)
	}
}
