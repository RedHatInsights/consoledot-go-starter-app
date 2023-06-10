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

// readinessProbe godoc
// @Summary      Determines readiness of the application
// @Description  Determines readiness of the application
// @Tags         probes
// @Produce      json
// @Success      200  {object}  readyResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /probes/ready [get]
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
// @Success      200  {object}  aliveResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /probes/alive [get]
func livelinessProbe(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"alive": "OK",
	})
}
