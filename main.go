package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
	"github.com/mimiiiz/timescale-testing/database"
	"github.com/mimiiiz/timescale-testing/handler"
	"github.com/mimiiiz/timescale-testing/model"
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

	conditionService := service.NewConditionService(db)
	conditionHandler := handler.NewConditionHandler(conditionService)
	conditionGroup := r.Group("conditions")

	conditionGroup.POST("/", conditionHandler.Create())
	conditionGroup.GET("/", conditionHandler.List())

	// r.Run()
	intervalInsertCondition(conditionService, locationService)

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

func intervalInsertCondition(c service.ConditionService, l service.LocationService) {
	rand.Seed(time.Now().UnixNano())
	for range time.Tick(30 * time.Second) {
		con := model.Condition{
			Time:        time.Now(),
			DeviceID:    randomdata.SillyName(),
			Temperature: randomdata.Decimal(-100, 100, 3),
			Humidity:    randomdata.Decimal(-50, 50, 3),
		}
		c.CreateCondition(&con)

		loc := model.Location{
			DeviceID:    randomdata.SillyName(),
			Location:    randomdata.Country(randomdata.FullCountry),
			Environment: randomdata.Street(),
		}
		l.CreateLocation(&loc)
	}
}
