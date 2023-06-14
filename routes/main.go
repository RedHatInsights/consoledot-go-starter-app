package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SetupRouter(apiPath string, db *pgx.Conn) *gin.Engine {
	router := gin.Default()
	setupProbes(router, apiPath)
	setupAPIRoutes(router, apiPath, db)
	return router
}
