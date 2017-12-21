package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mickaelmagniez/elastic-alert/controllers"
	"github.com/mickaelmagniez/elastic-alert/es"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	//SetupES()

	r := gin.Default()
	r.Use(CORSMiddleware())

	es.Init()

	alert := new(controllers.AlertController)

	r.POST("/alert", alert.Create)
	r.GET("/alert", alert.All)
	r.DELETE("/alert/:id", alert.Delete)
	r.Run()
}
