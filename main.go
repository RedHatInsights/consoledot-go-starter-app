package main

import (
	"os"

	"github.com/RedHatInsights/consoledot-go-starter-app/config"
	"github.com/RedHatInsights/consoledot-go-starter-app/docs"
	"github.com/RedHatInsights/consoledot-go-starter-app/providers/database"
	"github.com/RedHatInsights/consoledot-go-starter-app/routes"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	conf    = config.Load()
	db      = dbConnect()
	apiPath = makeAPIPath()
)

// main godoc
// @title ConsoleDot Go Starter App API
// @version         1.0
// @description     This is a sample API for the ConsoleDot Go Starter App.
// @contact.name   	Adam Drew
// @contact.url    	https://github.com/RedHatInsights/consoledot-go-starter-app
// @contact.email  	addrew@redhat.com
// @license.name  	MIT License
// @license.url   	https://opensource.org/license/mit/
// @BasePath  		/api/v1
func main() {
	router := routes.SetupRouter(apiPath)
	initAPIDocs(router)
	router.Run(conf.RouterBindAddress())
}

func makeAPIPath() string {
	return "/api/" + os.Getenv("API_PATH")
}

func dbConnect() *pgx.Conn {
	db, err := database.Connect(conf)
	if err != nil {
		panic(err)
	}
	return db
}

func initAPIDocs(router *gin.Engine) {
	// Example of editing the OpenAPI info programatically
	docs.SwaggerInfo.Host = conf.RouterBindAddress()
	docs.SwaggerInfo.BasePath = apiPath + "/v1"
	// Serve out the API Docs
	router.GET(apiPath, ginSwagger.WrapHandler(swaggerFiles.Handler))
}
