package aoc

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

func EnsureNoErrorOrPanic(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func ReadTextFileAllLines(path string) ([]string, error) {
	var content, err = ioutil.ReadFile(path)

	if err != nil {
		return nil, errors.New("File read failed!")
	}

	return strings.Split(string(content), "\n"), nil
}

func SliceConvertStringToInt(s []string) ([]int, error) {
	var nums = make([]int, 0, len(s))
	for _, v := range s {
		var val, e = strconv.Atoi(v)
		if e != nil {
			return nil, errors.New("Convert numbers failed!")
		}
		nums = append(nums, val)
	}
	return nums, nil
}
