package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func IndexHandler(context *gin.Context) {
	time := time.Now().Format("15:04:05")
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"time":      time,
		"questions": GetRandomQuestions(3),
	})
}
