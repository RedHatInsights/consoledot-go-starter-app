package routes

import (
	"context"
	"net/http"

	"github.com/RedHatInsights/consoledot-go-starter-app/metrics"
	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog/log"
)

func setupAPIRoutes(router *gin.Engine, apiPath string) {
	apiGroup := router.Group(apiPath + "/v1")
	addAPIRoutes(apiGroup)
}

func addAPIRoutes(apiGroup *gin.RouterGroup) {
	apiGroup.GET("/hello", helloWorld)
	apiGroup.GET("/db-info", dbInfo)
}

// helloWorld godoc
// @Summary      Recieve a greeting
// @Description  Recieve a greeting from the API
// @Tags         api
// @Produce      json
// @Success      200  {object}  map[string]any
// @Router       /api/starter-app-api/v1/hello [get]
func helloWorld(context *gin.Context) {
	metrics.IncrementRequests()
	context.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

// dbInfo godoc
// @Summary      Get database info
// @Description  Query the API to get some database informaiton
// @Tags         api
// @Produce      json
// @Success      200  {object}  map[string]any
// @Router       /api/starter-app-api/v1/db-info [get]
func dbInfo(ginContext *gin.Context) {
	// Guard against no database connection
	if !providerManager.DBProviderGuard() {
		log.Error().Msg("No database connection")
		metrics.IncrementErrors()
		ginContext.JSON(http.StatusInternalServerError, gin.H{
			"error": "No database connection",
		})
		return
	}

	var retVal string
	retStatus := http.StatusOK
	query := " select 'Database : ' ||current_database()||', '||'User : '|| user db_details;"
	err := providerManager.DBConnectionPool.QueryRow(context.Background(), query).Scan(&retVal)
	if err != nil {
		metrics.IncrementErrors()
		log.Error().Err(err).Msg("Error querying database")
		retVal = "Error querying database: " + err.Error()
		retStatus = http.StatusInternalServerError
	}
	metrics.IncrementRequests()

	ginContext.JSON(retStatus, gin.H{
		"db-info": retVal,
	})
}
