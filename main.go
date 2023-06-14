package main

import (
	"os"

	"github.com/RedHatInsights/consoledot-go-starter-app/config"
	"github.com/RedHatInsights/consoledot-go-starter-app/docs"
	"github.com/RedHatInsights/consoledot-go-starter-app/providers/database"
	"github.com/RedHatInsights/consoledot-go-starter-app/routes"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	conf     = config.Load()
	connPool *pgxpool.Pool
	apiPath  = makeAPIPath()
)

// main godoc
// @title ConsoleDot Go Starter App API
// @version         1.0
// @description     This is a sample API for the ConsoleDot Go Starter App.
// @contact.name   	Adam Drew
// @contact.email  	addrew@redhat.com
// @BasePath  		/api/starter-app-api/v1
func main() {
	connPool = dbConnect()
	defer connPool.Close()
	router := routes.SetupRouter(apiPath, connPool)
	initAPIDocs(router)
	router.Run(conf.RouterBindAddress())
}

func makeAPIPath() string {
	return "/api/" + os.Getenv("API_PATH")
}

func dbConnect() *pgxpool.Pool {
	connPool, err := database.Connect(conf)
	if err != nil {
		panic(err)
	}
	return connPool
}

func initAPIDocs(router *gin.Engine) {
	swaggerRoute := apiPath + "/api_docs/*any"
	// Example of editing the OpenAPI info programatically
	docs.SwaggerInfo.Host = conf.RouterBindAddress()
	docs.SwaggerInfo.BasePath = swaggerRoute
	// Serve out the API Docs

	router.GET(swaggerRoute, ginSwagger.WrapHandler(swaggerFiles.Handler))
}
