package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func setupAPIRoutes(router *gin.Engine, apiPath string, db *pgx.Conn) {
	apiGroup := router.Group(apiPath + "/v1")
	addAPIRoutes(apiGroup, db)
}

func addAPIRoutes(apiGroup *gin.RouterGroup, db *pgx.Conn) {
	apiGroup.GET("/hello", helloWorld)
	apiGroup.GET("/db-info", dbInfo(db))
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

// dbInfo godoc
// @Summary      Get database info
// @Description  Query the API to get some database informaiton
// @Tags         api
// @Produce      json
// @Success      200  {object}  map[string]any
// @Router       /api/starter-app/v1/db-info [get]
func dbInfo(db *pgx.Conn) func(context *gin.Context) {
	var retVal string
	retStatus := http.StatusOK
	query := " select 'Database : ' ||current_database()||', '||'User : '|| user db_details;"
	err := db.QueryRow(context.Background(), query).Scan(&retVal)
	if err != nil {
		retVal = "Error querying database: " + err.Error()
		retStatus = http.StatusInternalServerError
	}
	return func(context *gin.Context) {
		context.JSON(retStatus, gin.H{
			"db-info": retVal,
		})
	}
}
