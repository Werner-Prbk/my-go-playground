package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func loadNumbers(path string) ([]int, error) {
	var content, err = ioutil.ReadFile(path)

	if err != nil {
		return nil, errors.New("read file failed!")
	}

	var numStr = strings.Split(string(content), "\n")

	var nums = make([]int, 0, len(numStr))
	for _, v := range numStr {
		var val, e = strconv.Atoi(v)
		if e != nil {
			return nil, errors.New("Parse numbers failed!")
		}
		nums = append(nums, val)
	}
	return nums, nil
}

func sumOfSlice(s []int) (result int) {
	for _, v := range s {
		result += v
	}
	return result
}

func findNValsWhichSumToGivenValue(sum int, list []int, n int) ([]int, bool) {
	for i := 0; i < len(list)-n+1; i++ {
		if sumOfSlice(list[i:i+n]) == sum {
			return list[i : i+n], true
		}
	}

	return nil, false
}

func isSumOfTwoInList(sum int, list []int) bool {
	for i := 0; i < len(list); i++ {
		// shortcut
		if sum < list[i] {
			continue
		}

		for j := i; j < len(list); j++ {
			if list[i]+list[j] == sum {
				return true
			}
		}
	}

	return false
}

func findMin(s []int) (result int) {
	result = s[0]
	for _, v := range s {
		if result > v {
			result = v
		}
	}
	return result
}

func findMax(s []int) (result int) {
	result = s[0]
	for _, v := range s {
		if result < v {
			result = v
		}
	}
	return result
}

func main() {
	var numbers, err = loadNumbers("input.txt")

	if err != nil {
		panic("No valid input. Exit!")
	}

	var brokenNumIdx = 0

	for i := 25; i < len(numbers); i++ {
		if !isSumOfTwoInList(numbers[i], numbers[i-25:i]) {
			fmt.Printf("Found at line %v. Number is %v\n", i, numbers[i])
			brokenNumIdx = i
			break
		}
	}

	// part 2
	for n := 2; n <= brokenNumIdx; n++ {
		var nums, solution = findNValsWhichSumToGivenValue(numbers[brokenNumIdx], numbers[:brokenNumIdx], n)

		if solution == true {
			var solution = findMax(nums) + findMin(nums)
			fmt.Printf("The solution to break the encryption is %v\n", solution)
			break
		}
	}
}
