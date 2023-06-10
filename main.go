package main

import (
	"github.com/RedHatInsights/consoledot-go-starter-app/routes"
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
	router := routes.SetupRouter()
	router.Run()
}
