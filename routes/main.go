package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	setupProbesGroup(router)
	setupAPIRoutes(router)
	return router
}
