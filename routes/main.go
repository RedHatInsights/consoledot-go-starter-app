package routes

import (
	"github.com/RedHatInsights/consoledot-go-starter-app/providers"
	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
)

var (
	providerManager providers.Providers
)

func SetupRouter(apiPath string, pManager providers.Providers) *gin.Engine {
	providerManager = pManager
	router := gin.Default()
	router.Use(ginzerolog.Logger("gin"))
	setupProbes(router, apiPath)
	setupAPIRoutes(router, apiPath)
	return router
}
