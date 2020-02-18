package utils

import (
	"github.com/yanyiwu/gojieba"
)

var x *gojieba.Jieba

func init() {
	x = gojieba.NewJieba()
}

func CutWords(s string) []string {
	var words []string
	use_hmm := true
	words = x.Cut(s, use_hmm)
	return words
}
