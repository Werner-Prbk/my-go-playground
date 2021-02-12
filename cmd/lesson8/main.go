package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type operation string

const (
	acc operation = "acc"
	jmp operation = "jmp"
	nop operation = "nop"
)

type instruction struct {
	op  operation
	arg int
}

func parseInstruction(s string) instruction {
	var instr instruction
	instr.op = operation(s[:3])
	var arg, err = strconv.Atoi(s[5:])

	if err != nil {
		panic("cannot parse instruction!")
	}

	if s[4] == '+' {
		instr.arg = arg
		return instr
	}
	if s[4] == '-' {
		instr.arg = arg * -1
		return instr
	}

	panic("unexpected sign of instruction argument")
}

func readBootCode(path string) []instruction {
	var instr []instruction

	var content, e = ioutil.ReadFile(path)
	if e != nil {
		panic("Unexpected. Exit!")
	}

	for _, v := range strings.Split(string(content), "\n") {
		instr = append(instr, parseInstruction(v))
	}
	return instr
}

func executeUntilLoop(pc int, program []instruction) int {
	var marker []bool = make([]bool, len(program))
	var accu = 0

	for {
		var i = program[pc]

		// loop detected
		if marker[pc] == true {
			return accu
		}

		marker[pc] = true

		if i.op == acc {
			accu += i.arg
			pc++
		} else if i.op == jmp {
			pc += i.arg
		} else {
			pc++
		}
	}
}

func main() {
	var instr = readBootCode("input.txt")

	var accu = executeUntilLoop(0, instr)

	fmt.Printf("Accu is: %v\n", accu)
}
