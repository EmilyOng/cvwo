package controllers

import (
	"fmt"
	"net/http"

	"github.com/EmilyOng/cvwo/backend/models"
	stateService "github.com/EmilyOng/cvwo/backend/services/state"

	"github.com/gin-gonic/gin"
)

func CreateState(c *gin.Context) {
	var payload models.CreateStatePayload

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: error_UNEXPECTED})
		return
	}

	createStateResponse := stateService.CreateState(payload)
	c.JSON(http.StatusOK, createStateResponse)
}

func UpdateState(c *gin.Context) {
	var payload models.UpdateStatePayload

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.Response{Error: error_UNEXPECTED})
		return
	}

	updateStateResponse := stateService.UpdateState(payload)
	c.JSON(http.StatusOK, updateStateResponse)
}

func DeleteState(c *gin.Context) {
	var stateID uint8
	fmt.Sscan(c.Param("state_id"), &stateID)

	deleteStateResponse := stateService.DeleteState(models.DeleteStatePayload{ID: stateID})
	c.JSON(http.StatusOK, deleteStateResponse)
}
