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

func getYesCountOfGroup(answers []string) int {
	var m = make(map[byte]bool)

	for _, personAns := range answers {
		for _, v := range personAns {
			m[byte(v)] = true
		}
	}

	return len(m)
}

func main() {
	var answers = loadAnswers("answers.txt")

	var totalCnt = 0

	for _, v := range answers {
		totalCnt += getYesCountOfGroup(v)
	}

	fmt.Printf("Total sum of yes counts is %v\n", totalCnt)
}
