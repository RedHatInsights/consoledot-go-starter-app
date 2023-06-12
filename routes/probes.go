package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupProbes(router *gin.Engine, apiPath string) {
	router.GET("/readyz", readinessProbe)
	router.GET("/healthz", livelinessProbe)
}

// readinessProbe godoc
// @Summary      Determines readiness of the application
// @Description  Determines readiness of the application
// @Tags         probes
// @Produce      json
// @Success      200  {object}  map[string]any
// @Router       /healthz [get]
func readinessProbe(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"ready": "OK",
	})
}

// livelinessProbe godoc
// @Summary      Determines if application is still alive
// @Description  Determines if application is still alive
// @Tags         probes
// @Produce      json
// @Success      200  {object}  map[string]any
// @Router       /livez [get]
func livelinessProbe(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"alive": "OK",
	})
}
