package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const tree byte = '#'
const square byte = '.'

func loadMap(path string) [][]byte {
	var data, err = ioutil.ReadFile(path)

	if err != nil {
		panic("This is unexpected")
	}

	var res [][]byte
	for _, v := range strings.Split(string(data), "\n") {
		res = append(res, []byte(v))
	}
	return res
}

func isTreeAtPos(theMap [][]byte, x, y int) bool {
	x = x % len(theMap[0])
	return theMap[y][x] == tree
}

func countTrees(theMap [][]byte, toRight int, toBottom int) int {
	var cnt = 0
	var x = 0
	for row := 0; row < len(theMap); row += toBottom {
		if isTreeAtPos(theMap, x, row) {
			cnt++
		}
		x += toRight
	}
	return cnt
}

func main() {
	var theMap = loadMap("input.txt")
	var treeCnt = countTrees(theMap, 3, 1)

	fmt.Printf("Found trees %v\n", treeCnt)
}
