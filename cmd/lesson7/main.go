package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func loadRuleList(path string) []string {
	var content, err = ioutil.ReadFile(path)

	if err != nil {
		panic("Unexpected error!")
	}

	return strings.Split(string(content), "\n")
}

const MyBagColor string = "shiny gold"

type bagColorRule struct {
	bagColor   string
	canContain map[string]int
}

func parseCanContain(canContain string) (string, int) {
	canContain = strings.Trim(canContain, " .")

	if canContain[:3] == "no " {
		return "", 0
	} else {
		var numStr = strings.Split(canContain, " ")[0]
		var num, _ = strconv.Atoi(numStr)
		return canContain[len(numStr)+1:], num
	}
}

func parseRule(rule string) bagColorRule {
	// cleanup the stuff with no info
	var cleaned = strings.ReplaceAll(rule, "bags", "")
	cleaned = strings.ReplaceAll(cleaned, "bag", "")

	var splitted = strings.Split(cleaned, "contain")
	var what = splitted[0]
	var contains = strings.Split(splitted[1], ",")

	var bc = bagColorRule{
		bagColor:   strings.TrimSpace(what),
		canContain: make(map[string]int, len(contains)),
	}

	for _, v := range contains {
		var bagColor, count = parseCanContain(v)
		bc.canContain[bagColor] = count
	}
	return bc
}

func containsMyBag(bcr bagColorRule, allBcr map[string]bagColorRule, myBag string) bool {
	if bcr.bagColor == myBag {
		return true
	}

	for k, v := range bcr.canContain {
		if v != 0 {
			if containsMyBag(allBcr[k], allBcr, myBag) {
				return true
			}
		}
	}

	return false
}

func main() {
	var rulesStr = loadRuleList("rules.txt")

	var rules = make(map[string]bagColorRule, len(rulesStr))

	for _, v := range rulesStr {
		var r = parseRule(v)
		rules[r.bagColor] = r
	}

	var cnt = 0

	// find number of bags which may contain my bag
	for _, v := range rules {
		if v.bagColor != MyBagColor {
			if containsMyBag(v, rules, MyBagColor) {
				cnt++
			}
		}
	}

	fmt.Printf("%v bags may contain my bag\n", cnt)
}
