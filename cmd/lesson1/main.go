package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func loadNumbers(path string) []int {
	var data, err = ioutil.ReadFile(path)

	if err != nil {
		panic("This is unexpected")
	}

	var numbStr = strings.Split(string(data), "\n")
	var numbInt = make([]int, len(numbStr))

	for i := 0; i < len(numbStr); i++ {
		numbInt[i], _ = strconv.Atoi(numbStr[i])
	}

	return numbInt
}

func getNumbersWithSum(numbers []int, sum int) (int, int, error) {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == sum {
				return numbers[i], numbers[j], nil
			}
		}
	}
	return 0, 0, errors.New("No solution found!")
}

func main() {
	var numbers = loadNumbers("numbers.txt")

	var n1, n2, err = getNumbersWithSum(numbers, 2020)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("The found numbers are %v and %v. The solution is %v!\n", n1, n2, n1*n2)
	}
}
