package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SetupRouter(apiPath string, connPool *pgxpool.Pool) *gin.Engine {
	router := gin.Default()
	setupProbes(router, apiPath)
	setupAPIRoutes(router, apiPath, connPool)
	return router
}
