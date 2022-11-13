package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (service *Controller) HealthCheck(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{"status": "ok"})

}
