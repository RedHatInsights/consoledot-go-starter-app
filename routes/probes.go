package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupProbesGroup(router *gin.Engine) {
	probesGroup := router.Group("/probes")
	addProbesRoutes(probesGroup)
}

func addProbesRoutes(probesGroup *gin.RouterGroup) {
	probesGroup.GET("/ready", readinessProbe)
	probesGroup.GET("/alive", livelinessProbe)
}

func readinessProbe(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"ready": "OK",
	})
}

func livelinessProbe(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"alive": "OK",
	})
}
