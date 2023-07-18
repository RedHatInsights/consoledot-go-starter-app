package main

import (
	"os"

	"github.com/RedHatInsights/consoledot-go-starter-app/config"
	"github.com/RedHatInsights/consoledot-go-starter-app/docs"
	"github.com/RedHatInsights/consoledot-go-starter-app/metrics"
	"github.com/RedHatInsights/consoledot-go-starter-app/providers/database"
	"github.com/RedHatInsights/consoledot-go-starter-app/routes"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	conf     = config.Load()
	connPool database.ConnectionPool
	apiPath  = conf.GetApiPath()
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
	// Create the database connection pool and defer closing it
	connPool = dbConnect()
	defer connPool.Close()
	// Serve the prometheus metrics
	go metrics.Serve(conf)
	// Set up the Gin router
	router := routes.SetupRouter(apiPath, connPool)
	// Set up the OpenAPI docs
	initAPIDocs(router)
	// Run the router
	router.Run(conf.RouterBindAddress())
}

// initLogging sets up the logging
func initLogging() {
	// Set the default log level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("Starting ConsoleDot Go Starter App")
}

// dbConnect returns the datbase connection pool
func dbConnect() database.ConnectionPool {
	connPool, err := database.Connect(conf)
	if err != nil {
		log.Error().Err(err).Msg("Error connecting to database. Exiting ...")
		os.Exit(1)
	}
	return connPool
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
