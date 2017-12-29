package router

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/mickaelmagniez/elastic-alert/api/controllers"
	"github.com/mickaelmagniez/elastic-alert/api/router/middleware"
)

func Load() http.Handler {
	r := gin.New()
	r.Use(middleware.CORS())

	alerts := r.Group("/alerts")
	{
		alertsController := new(controllers.AlertsController)

		alerts.POST("", alertsController.Create)
		alerts.GET("", alertsController.All)
		alerts.GET("/:id", alertsController.Get)
		alerts.DELETE("/:id", alertsController.Delete)
		alerts.PUT("/:id", alertsController.Update)
	}

	elastics := r.Group("/elastics")
	{
		elasticsController := new(controllers.ElasticsController)

		elastics.GET("", elasticsController.GetServers)
		elastics.GET("/indices", elasticsController.GetIndices)
		elastics.GET("/types", elasticsController.GetTypes)
	}

	r.Run()

	return r
}
