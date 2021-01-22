package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type PassportKey string

const (
	BirthYear      PassportKey = "byr"
	IssueYear      PassportKey = "iyr"
	ExpirationYear PassportKey = "eyr"
	Height         PassportKey = "hgt"
	HairColor      PassportKey = "hcl"
	EyeColor       PassportKey = "ecl"
	PassportId     PassportKey = "pid"
	CountryId      PassportKey = "cid"
)

type Passport map[PassportKey]string
type PassportKeyValidator func(value string) bool
type PassportValidator map[PassportKey]PassportKeyValidator

func parsePassport(s string) (p Passport) {
	p = make(Passport)

	s = strings.Replace(s, "\n", " ", -1)

	for _, v := range strings.Split(s, " ") {
		var kv = strings.Split(v, ":")
		p[PassportKey(kv[0])] = kv[1]
	}

	return p
}

func readBatch(path string) (passports []Passport) {
	var data, err = ioutil.ReadFile(path)

	if err != nil {
		panic("This is unexpected")
	}

	// separated via empty line
	var pp = strings.Split(string(data), "\n\n")

	for _, v := range pp {
		passports = append(passports, parsePassport(v))
	}
	return passports
}

func validatePassport(passport Passport, validator PassportValidator) bool {
	for entry, validator := range validator {
		var value, ok = passport[entry]

		// passport must contain the entry and the
		// the specific entry must be valid
		if !ok || !validator(value) {
			return false
		}
	}
	return true
}

func main() {
	var passports = readBatch("batch.txt")
	var validCnt = 0

	var passportWithReqiredKeys = PassportValidator{
		BirthYear: func(x string) bool {
			return isIntegerWithinRange(x, 1920, 2002)
		},
		ExpirationYear: func(x string) bool {
			return isIntegerWithinRange(x, 2020, 2030)
		},
		EyeColor:  isEyeColor,
		HairColor: isHairColor,
		Height:    isHeight,
		IssueYear: func(x string) bool {
			return isIntegerWithinRange(x, 2010, 2020)
		},
		PassportId: isPassportId,
	}

	for _, v := range passports {
		if validatePassport(v, passportWithReqiredKeys) {
			validCnt++
		}
	}

	fmt.Printf("Number of valid passports is %v out of %v\n", validCnt, len(passports))
}

func isEyeColor(s string) bool {
	if strings.Compare("amb", s) == 0 ||
		strings.Compare("blu", s) == 0 ||
		strings.Compare("brn", s) == 0 ||
		strings.Compare("gry", s) == 0 ||
		strings.Compare("grn", s) == 0 ||
		strings.Compare("hzl", s) == 0 ||
		strings.Compare("oth", s) == 0 {
		return true
	}
	return false
}

func isIntegerWithinRange(s string, min, max int) bool {
	var val, err = strconv.Atoi(s)
	if err == nil && val >= min && val <= max {
		return true
	}
	return false
}

func isHairColor(s string) bool {
	if len(s) != 7 {
		return false
	}
	if s[0] != '#' {
		return false
	}
	for i := 1; i < len(s); i++ {
		if !((s[i] >= '0' && s[i] <= '9') || (s[i] >= 'a' && s[i] <= 'f')) {
			return false
		}
	}
	return true
}

func isHeight(s string) bool {
	var unit = s[len(s)-2:]
	var valueStr = s[:len(s)-2]
	var value, err = strconv.Atoi(valueStr)

	if err != nil {
		return false
	}

	if unit == "cm" {
		if value >= 150 && value <= 193 {
			return true
		}
	} else if unit == "in" {
		if value >= 59 && value <= 76 {
			return true
		}
	}
	return false
}

func isPassportId(s string) bool {
	if len(s) != 9 {
		return false
	}

	for _, v := range s {
		if !(byte(v) >= '0' && byte(v) <= '9') {
			return false
		}
	}
	return true
}
