package routes

import (
	"github.com/RedHatInsights/consoledot-go-starter-app/providers/database"
	"github.com/gin-gonic/gin"
)

func SetupRouter(apiPath string, connPool database.ConnectionPool) *gin.Engine {
	router := gin.Default()
	setupProbes(router, apiPath)
	setupAPIRoutes(router, apiPath, connPool)
	return router
}
