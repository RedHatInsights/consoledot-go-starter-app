package main

import (
	"github.com/RedHatInsights/consoledot-go-starter-app/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run()
}
