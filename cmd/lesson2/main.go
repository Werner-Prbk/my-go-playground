package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type passwordEntry struct {
	min      int
	max      int
	char     byte
	password string
}

func parseLine(line string) passwordEntry {
	var entry passwordEntry

	// format: "{min}-{max} {char}: {password}""

	var three = strings.Split(line, " ")
	var minMax = strings.Split(three[0], "-")

	entry.min, _ = strconv.Atoi(minMax[0])
	entry.max, _ = strconv.Atoi(minMax[1])
	entry.char = three[1][0]
	entry.password = three[2]
	return entry
}

func loadPasswordEntries(path string) []passwordEntry {
	var data, err = ioutil.ReadFile(path)

	if err != nil {
		panic("This is unexpected")
	}

	var pweList []passwordEntry

	for _, v := range strings.Split(string(data), "\n") {
		pweList = append(pweList, parseLine(v))
	}

	return pweList
}

func validatePasswordV1(e passwordEntry) bool {
	var cnt = 0
	for _, v := range e.password {
		if byte(v) == e.char {
			cnt++
		}
	}

	return cnt >= e.min && cnt <= e.max
}

func validatePasswordV2(e passwordEntry) bool {
	if len(e.password) < e.max {
		return false
	}

	var cnt = 0
	if byte(e.password[e.min-1]) == e.char {
		cnt++
	}
	if byte(e.password[e.max-1]) == e.char {
		cnt++
	}

	return cnt == 1
}

func main() {
	var pweList = loadPasswordEntries("passwords.txt")

	var validCnt = 0
	for _, v := range pweList {
		if validatePasswordV2(v) {
			validCnt++
		}
	}

	fmt.Printf("The number of valid passwords is %d out of %d.\n", validCnt, len(pweList))
}
