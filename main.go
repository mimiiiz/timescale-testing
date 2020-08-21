package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mimiiiz/timescale-testing/database"
	"github.com/mimiiiz/timescale-testing/handler"
	"github.com/mimiiiz/timescale-testing/service"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	r := gin.Default()
	r.Use(corsMiddleware())

	locationService := service.NewLocationService(db)
	locationHandler := handler.NewLocationHandler(locationService)
	locationGroup := r.Group("locations")

	locationGroup.POST("/", locationHandler.Create())
	locationGroup.GET("/", locationHandler.List())

	r.Run()

}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "true")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
