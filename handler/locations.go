package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mimiiiz/timescale-testing/model"
	"github.com/mimiiiz/timescale-testing/service"
)

type locationHandler struct {
	locationService service.LocationService
}

type LocationHandler interface {
	Create() gin.HandlerFunc
	List() gin.HandlerFunc
}

func NewLocationHandler(locationService service.LocationService) LocationHandler {
	return &locationHandler{
		locationService,
	}
}

func (h *locationHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var location model.Location
		err := c.ShouldBindJSON(&location)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if err := h.locationService.CreateLocation(&location); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, location)
	}
}

func (h *locationHandler) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, _ := h.locationService.ListLocation()
		c.JSON(http.StatusOK, users)
	}
}
