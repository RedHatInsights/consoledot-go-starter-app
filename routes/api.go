package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupAPIRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api/v1")
	addAPIRoutes(apiGroup)
}

func addAPIRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/hello", helloWorld)
}

// helloWorld godoc
// @Summary      Recieve a greeting
// @Description  Recieve a greeting from the API
// @Tags         api
// @Produce      json
// @Success      200  {object}  map[string]any
// @Router       /api/v1/hello [get]
func helloWorld(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
