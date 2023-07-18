package main

import (
	"github.com/RedHatInsights/consoledot-go-starter-app/config"
	"github.com/RedHatInsights/consoledot-go-starter-app/docs"
	"github.com/RedHatInsights/consoledot-go-starter-app/metrics"
	"github.com/RedHatInsights/consoledot-go-starter-app/providers"
	"github.com/RedHatInsights/consoledot-go-starter-app/routes"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	conf            = config.Load()
	providerManager providers.Providers
	apiPath         = conf.GetApiPath()
)

// main godoc
// @title ConsoleDot Go Starter App API
// @version         1.0
// @description     This is a sample API for the ConsoleDot Go Starter App.
// @contact.name   	Adam Drew
// @contact.email  	addrew@redhat.com
// @BasePath  		/api/starter-app-api/v1
func main() {
	// Initialize logging
	initLogging()

	var closeFunc func()
	providerManager, closeFunc = initProviders()
	defer closeFunc()

	// Serve the prometheus metrics
	go metrics.Serve(conf)
	// Set up the Gin router
	router := routes.SetupRouter(apiPath, providerManager)
	// Set up the OpenAPI docs
	initAPIDocs(router)
	// Run the router
	router.Run(conf.RouterBindAddress())
}

func initProviders() (providers.Providers, func()) {
	p := providers.Init(conf)
	return p, func() {
		if p.DBConnectionPool != nil {
			p.DBConnectionPool.Close()
		}
	}
}

// initLogging sets up the logging
func initLogging() {
	// Set the default log level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("Starting ConsoleDot Go Starter App")
}

// initAPIDocs sets up the swagger (openAPI) docs
func initAPIDocs(router *gin.Engine) {
	swaggerRoute := apiPath + "/api_docs/*any"
	// Example of editing the OpenAPI info programatically
	docs.SwaggerInfo.Host = conf.RouterBindAddress()
	docs.SwaggerInfo.BasePath = swaggerRoute
	// Serve out the API Docs
	router.GET(swaggerRoute, ginSwagger.WrapHandler(swaggerFiles.Handler))
}
