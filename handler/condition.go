package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mimiiiz/timescale-testing/model"
	"github.com/mimiiiz/timescale-testing/service"
)

type conditionHandler struct {
	conditionService service.ConditionService
}

type ConditionHandler interface {
	Create() gin.HandlerFunc
	List() gin.HandlerFunc
}

func NewConditionHandler(conditionService service.ConditionService) ConditionHandler {
	return &conditionHandler{
		conditionService,
	}
}

func (h *conditionHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var condition model.Condition
		err := c.ShouldBindJSON(&condition)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if err := h.conditionService.CreateCondition(&condition); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, condition)
	}
}

func (h *conditionHandler) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, _ := h.conditionService.ListCondition()
		c.JSON(http.StatusOK, users)
	}
}
