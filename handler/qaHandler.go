package handler

import (
	"XgfyQA/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var qs []string
var dict map[string]interface{}
var cutWords [][]string

func init() {
	log.Println("Init")
	qs, dict = utils.ReadJson("./data.json")
	cutWords = make([][]string, len(qs))
	for i, q := range qs {
		words := utils.CutWords(q)
		cutWords[i] = words
	}
}

func QuestionHandler(context *gin.Context) {
	var answer interface{}
	question := context.Query("question")
	score := make([]float64, len(qs))
	srcWords := utils.CutWords(question)
	for i, distWords := range cutWords {
		s := utils.CosineSimilar(srcWords, distWords)
		score[i] = s
	}
	maxIndex, maxScore := utils.MaxIndex(score)
	log.Println(maxIndex, maxScore)
	if maxScore < 0.5 {
		answer = "暂时还不能回答你的问题"
	} else {
		findQuestion := qs[maxIndex]
		answer = dict[findQuestion]
	}
	// 回答太快了，没有感觉
	time.Sleep(500 * time.Millisecond)
	//返回json
	context.JSON(http.StatusOK, gin.H{
		"data":     answer,
		"message":  "success",
		"question": question,
	})
}
