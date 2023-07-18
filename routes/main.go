package routes

import (
	"github.com/RedHatInsights/consoledot-go-starter-app/providers/database"
	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
)

func SetupRouter(apiPath string, connPool database.ConnectionPool) *gin.Engine {
	router := gin.Default()
	router.Use(ginzerolog.Logger("gin"))
	setupProbes(router, apiPath)
	setupAPIRoutes(router, apiPath, connPool)
	return router
}
