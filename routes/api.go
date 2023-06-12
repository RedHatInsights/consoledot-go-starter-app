package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupAPIRoutes(router *gin.Engine, apiPath string) {
	apiGroup := router.Group(apiPath + "/v1")
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
// @Router       /api/starter-app/v1/hello [get]
func helloWorld(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
