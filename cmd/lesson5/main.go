package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readBoardingPasses(path string) []string {
	var content, err = ioutil.ReadFile(path)
	if err != nil {
		panic("Cannot read input")
	}
	return strings.Split(string(content), "\n")
}

type BoardingPass struct {
	Row int
	Col int
	Id  int
}

func recursiveFind(low, high int, pattern string, match byte) int {
	if low == high {
		return low
	}

	var totalLeft = 1 + high - low

	if pattern[0] == match {
		return recursiveFind(low, low+totalLeft/2-1, pattern[1:], match)
	}
	return recursiveFind(low+totalLeft/2, high, pattern[1:], match)
}

func findRow(rowString string) int {
	return recursiveFind(0, 127, rowString, 'F')
}

func findCol(colString string) int {
	return recursiveFind(0, 7, colString, 'L')
}

func parseBoardingPass(bpString string) BoardingPass {
	var bp BoardingPass
	bp.Row = findRow(bpString[0:7])
	bp.Col = findCol(bpString[7:10])
	bp.Id = bp.Row*8 + bp.Col
	return bp
}

func main() {
	var bpRaw = readBoardingPasses("boardingpasses.txt")

	var boardingPasses = make(map[int]BoardingPass, len(bpRaw))

	for _, v := range bpRaw {
		var parsed = parseBoardingPass(v)

		// check if same id is existing already
		var _, ok = boardingPasses[parsed.Id]

		if ok {
			fmt.Println("ID is not unique. There is an error!!!")
			return
		}
		boardingPasses[parsed.Id] = parsed
	}

	// find max seat id
	var max = 0
	for k, _ := range boardingPasses {
		if k > max {
			max = k
		}
	}

	fmt.Printf("The highest seat id is %v\n", max)
}
