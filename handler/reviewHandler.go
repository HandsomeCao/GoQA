package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var reviewDict map[string]string
var answers map[string]string

func init() {
	reviewDict = make(map[string]string)
	reviewDict["review1"] = "请输入你的姓名。"
	reviewDict["review2"] = "请输入你现在所在的地点。"
	reviewDict["review3"] = "您最近14天内是否到过或经停武汉？请回答\"是\"或者\"否\"。"
	reviewDict["review4"] = "您最近14天内是否出现发热、乏力、干咳、呼吸困难等症状？请回答\"是\"或者\"否\"。"
	reviewDict["review5"] = "您最近14天内是否与来自武汉市或者周边地区的人员有过较为亲密的接触？请回答\"是\"或者\"否\"。"
	reviewDict["review6"] = "您最近14天内是否接触或食用过野味、去过野味市场或野生动物栖息地？请回答\"是\"或者\"否\"。"
	reviewDict["review7"] = "是否有任何疫情相关的，需要注意的情况？请回答\"是\"或者\"否\"。"
	reviewDict["review8"] = "请补充您情况的描述信息，若无请回答\"无\"。"
	reviewDict["review9"] = ""
	answers = make(map[string]string, 8)
}

// ReviewHandler recive review request
func ReviewHandler(context *gin.Context) {
	responseJudge := [5]string{"review4", "review5", "review6", "review7", "review8"}
	question := context.Query("question")
	reviewID := context.Query("id")
	final := "根据您的回答，您暂时安全，但一定要戴好口罩，尽量不要出门，祝您健康！"
	time.Sleep(500 * time.Millisecond)
	if v, ok := reviewDict[reviewID]; ok {
		correct := true
		if reviewID != "review9" {
			// 对回答进行判断是否为“是”或“不是”
			for _, i := range responseJudge {
				if reviewID == i && question != "是" && question != "否" {
					v = "答案必须为\"是\"或\"否\"，请重新回答。"
					correct = false
					break
				}
			}
			// 若满足回答条件
			context.JSON(http.StatusOK, gin.H{
				"data":      v,
				"message":   "success",
				"question":  question,
				"recommand": "",
				"isCorrect": correct,
			})
			if correct {
				answers[reviewID] = question
			}
		} else {
			log.Println(answers)
			for _, v := range answers {
				if v == "是" {
					final = "根据您的回答，您可能有染上肺炎的风险，建议您去当地医院检查，祝您健康！"
					break
				}
			}
			recommand_questions := GetRandomQuestions(3)
			context.JSON(http.StatusOK, gin.H{
				"data":      final,
				"message":   "success",
				"question":  question,
				"recommand": recommand_questions,
				"isCorrect": true,
			})
			answers = make(map[string]string, 8)
		}
	} else {
		log.Panicln("未找到相应问题。")
	}
}
