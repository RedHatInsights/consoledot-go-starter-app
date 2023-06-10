package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupAPIRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	addAPIRoutes(apiGroup)
}

func addAPIRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/hello", helloWorld)
}

func helloWorld(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
