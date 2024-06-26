package router

import (
	"os"
	"webCrawler/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	mainDir, _ := os.Getwd()

	router.Static("/static/css", mainDir+"/static/css")

	router.LoadHTMLGlob(mainDir + "/templates/*")

	router.GET("/search", handlers.SearchHandelerForWebCrawler)
	router.POST("/search", handlers.SearchHandelerForWebCrawler)

	return router
}
