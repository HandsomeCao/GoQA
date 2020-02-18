package utils

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
)

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
	for k, _ := range m {
		keys[j] = k
		j++
	}
	return keys
}

func MaxIndex(score []float64) (int, float64) {
	max, index := 0.0, 0
	for i, v := range score {
		if v > max {
			max = v
			index = i
		}
	}
	return index, max
}
