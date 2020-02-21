package handler

import (
	"XgfyQA/utils"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
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
	var question string
	question = context.Query("question")
	score := make([]float64, len(qs))
	srcWords := utils.CutWords(question)
	for i, distWords := range cutWords {
		s := utils.CosineSimilar(srcWords, distWords)
		score[i] = s
	}
	topIndexs, topScores := utils.MaxIndex(score, 4)
	maxScore, maxIndex := topScores[0], topIndexs[0]
	//log.Println("分数是:", topIndexs, topScores)
	if maxScore < 0.3 {
		answer = "暂时还不能回答你的问题"
	} else {
		findQuestion := qs[maxIndex]
		answer = dict[findQuestion]
	}
	// 回答太快了，没有感觉
	time.Sleep(500 * time.Millisecond)
	// 获取推荐问题
	recommandString := make([]string, 3)
	for i := 1; i < 4; i++ {
		question = qs[topIndexs[i]]
		recommandString[i-1] = question
	}
	//返回json
	context.JSON(http.StatusOK, gin.H{
		"data":      answer,
		"message":   "success",
		"question":  question,
		"recommand": recommandString,
	})
}

func GetRandomQuestions(num int) []string {
	length := len(qs)
	questions := make([]string, num)
	for i := 0; i < num; i++ {
		id := rand.Intn(length)
		questions[i] = qs[id]
	}
	return questions
}
