package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const tree byte = '#'
const square byte = '.'

type map2d []string

func loadMap(path string) map2d {
	var data, err = ioutil.ReadFile(path)

	if err != nil {
		panic("This is unexpected")
	}
	return strings.Split(string(data), "\n")
}

func isTreeAtPos(theMap map2d, x, y int) bool {
	x = x % len(theMap[0])
	return theMap[y][x] == tree
}

func countTrees(theMap map2d, toRight int, toBottom int) int {
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
	var solution = countTrees(theMap, 1, 1)
	solution *= countTrees(theMap, 3, 1)
	solution *= countTrees(theMap, 5, 1)
	solution *= countTrees(theMap, 7, 1)
	solution *= countTrees(theMap, 1, 2)

	fmt.Printf("The solution is %v\n", solution)
}
