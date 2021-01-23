package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func loadAnswers(path string) [][]string {
	var content, err = ioutil.ReadFile(path)

	if err != nil {
		panic("This is unexpected")
	}

	var groups = strings.Split(string(content), "\n\n")
	var answers [][]string

	for _, v := range groups {
		answers = append(answers, strings.Split(v, "\n"))
	}
	return answers
}

func getAnyYesCountOfGroup(answers []string) int {
	var m = make(map[byte]bool)

	for _, personAns := range answers {
		for _, v := range personAns {
			m[byte(v)] = true
		}
	}

	return len(m)
}

func getAllYesCountOfGroup(answers []string) int {
	var m = make(map[byte]int)

	for _, personAns := range answers {
		for _, v := range personAns {
			var val, ok = m[byte(v)]
			if ok {
				m[byte(v)] = val + 1
			} else {
				m[byte(v)] = 1
			}
		}
	}

	var expectedYesCnt = len(answers)
	var result = 0

	for _, v := range m {
		if v == expectedYesCnt {
			result++
		}
	}
	return result
}

func main() {
	var answers = loadAnswers("answers.txt")

	var totalAnyYesCnt = 0
	var totalAllYesCnt = 0

	for _, v := range answers {
		totalAnyYesCnt += getAnyYesCountOfGroup(v)
		totalAllYesCnt += getAllYesCountOfGroup(v)
	}

	fmt.Printf("Total sum of any yes counts is %v\n", totalAnyYesCnt)
	fmt.Printf("Total sum of all yes counts is %v\n", totalAllYesCnt)
}
