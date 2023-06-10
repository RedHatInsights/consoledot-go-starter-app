package main

import (
	"github.com/RedHatInsights/consoledot-go-starter-app/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/RedHatInsights/consoledot-go-starter-app/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// main godoc
// @title           ConsoleDot Go Starter App API
// @version         1.0
// @description     This is a sample API for the ConsoleDot Go Starter App.
// @contact.name   	Adam Drew
// @contact.url    	https://github.com/RedHatInsights/consoledot-go-starter-app
// @contact.email  	addrew@redhat.com
// @license.name  	MIT License
// @license.url   	https://opensource.org/license/mit/
// @host      		localhost:8080
// @BasePath  		/api/v1
func main() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	router := routes.SetupRouter()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
