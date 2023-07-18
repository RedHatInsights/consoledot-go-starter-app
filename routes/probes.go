package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupProbes(router *gin.Engine, apiPath string) {
	router.GET("/readyz", readinessProbe)
	router.GET("/livez", livelinessProbe)
}

// readinessProbe godoc
// @Summary      Determines readiness of the application
// @Description  Determines readiness of the application
// @Tags         probes
// @Produce      json
// @Success      200  {object}  map[string]any
// @Router       /readyz [get]
func readinessProbe(context *gin.Context) {
	// Your probes should not just return OK, but should also return useful information about the state of your
	// application. Reporting on the state of dependencies is a good example of this. However, this must be
	// balanced against the impact of the probes being hit frequently by the kubernetes API. The right
	// balance will depend on your application and your environment.
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
// @Router       /healthz [get]
func livelinessProbe(context *gin.Context) {
	// Your probes should not just return OK, but should also return useful information about the state of your
	// application. Reporting on the state of dependencies is a good example of this. However, this must be
	// balanced against the impact of the probes being hit frequently by the kubernetes API. The right
	// balance will depend on your application and your environment.
	context.JSON(http.StatusOK, gin.H{
		"alive": "OK",
	})
}
