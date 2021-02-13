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

func executeUntilLoopOrEnd(pc int, program []instruction) (int, bool) {
	var marker []bool = make([]bool, len(program))
	var accu = 0

	for {
		var i = program[pc]

		// loop detected
		if marker[pc] == true {
			return accu, false
		}
		// end detected
		if pc == (len(program) - 1) {
			return accu, true
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

	var accu, _ = executeUntilLoopOrEnd(0, instr)

	fmt.Printf("Accu is: %v\n", accu)

	// try to fix "bug" in instructions
	for i := 0; i < len(instr); i++ {
		var opAtPc = instr[i].op
		if opAtPc == jmp {
			instr[i].op = nop
		} else if (opAtPc == nop) && (instr[i].arg != 0) {
			instr[i].op = jmp
		} else {
			// nothing changed, dont waste time...
			continue
		}

		var accu, endReached = executeUntilLoopOrEnd(0, instr)

		if endReached {
			fmt.Printf("Modified line %v from %v to %v. Accu: %v\n", i, opAtPc, instr[i].op, accu)
			break
		} else {
			// undo modification
			instr[i].op = opAtPc
		}
	}
}
