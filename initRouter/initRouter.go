// Package initRouter provides ...
package initRouter

import (
	"XgfyQA/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// use default Engine with default middleware Logger and Recover
	router := gin.Default()
	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	// load static files
	router.Static("/statics", "./statics")

	index := router.Group("/")
	{
		index.GET("", handler.IndexHandler)
	}
	questionRouter := router.Group("/question")
	{
		questionRouter.GET("", handler.QuestionHandler)
	}
	return router
}
