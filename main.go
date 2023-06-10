package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloWorld(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

func readinessProbe(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"ready": "OK",
	})
}

func livelinessProbe(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"alive": "OK",
	})
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/api/hello", helloWorld)
	router.GET("/probes/ready", readinessProbe)
	router.GET("/probes/alive", livelinessProbe)
	return router
}

func main() {
	router := setupRouter()
	router.Run()
}
