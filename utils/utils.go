package utils

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"sort"
)

type Slice struct {
	value []float64
	idx   []int
}

func (s Slice) Swap(i, j int) {
	s.value[i], s.value[j] = s.value[j], s.value[i]
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func (s Slice) Len() int {
	return len(s.value)
}

func (s Slice) Less(i, j int) bool {
	return s.value[j] < s.value[i]
}

func readFile(path string) string {
	// open file
	filePtr, err := os.Open(path)
	if err != nil {
		log.Println("Open file failed ", err.Error())
	}
	defer filePtr.Close()

	// json process
	var s string
	inputReader := bufio.NewReader(filePtr)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break
		}
		s = s + inputString
	}
	return s
}

func JSONToMap(str string) map[string]interface{} {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		log.Panicln(err.Error())
	}
	return tempMap
}

func ReadJson(path string) ([]string, map[string]interface{}) {
	jsonString := readFile(path)
	dict := JSONToMap(jsonString)
	return getKey(dict), dict
}

func getKey(m map[string]interface{}) []string {
	//先定义数组
	j := 0
	keys := make([]string, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

// MaxIndex get sort indices
func MaxIndex(score []float64, topK int) ([]int, []float64) {
	idx := make([]int, len(score))
	for i:= 0; i< len(score); i++{
		idx[i] = i
	}
	scoreSlice := Slice{
		value: score,
		idx: idx,
	}
	sort.Sort(scoreSlice)
	// for i, v := range score {
	// 	if v > max {
	// 		max = v
	// 		index = i
	// 	}
	// }
	return scoreSlice.idx[:topK], scoreSlice.value[:topK]
}
