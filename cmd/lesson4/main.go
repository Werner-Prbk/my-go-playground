package main

import (
	"fmt"
	"io/ioutil"
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

func validatePassport(passport Passport, validator Passport) bool {
	for k := range validator {
		var _, ok = passport[k]
		if !ok {
			return false
		}
	}
	return true
}

func main() {
	var passports = readBatch("batch.txt")
	var validCnt = 0

	var passportWithReqiredKeys = Passport{
		BirthYear:      "",
		ExpirationYear: "",
		EyeColor:       "",
		HairColor:      "",
		Height:         "",
		IssueYear:      "",
		PassportId:     "",
	}

	for _, v := range passports {
		if validatePassport(v, passportWithReqiredKeys) {
			validCnt++
		}
	}

	fmt.Printf("Number of valid passports is %v out of %v\n", validCnt, len(passports))
}
