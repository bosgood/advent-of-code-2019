package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func try(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("inputs/day2_part1.txt")
	try(err)
	b, err := ioutil.ReadAll(f)
	try(err)
	trimmed := bytes.TrimSpace(b)
	fmt.Println(Day2Part1(string(trimmed)))
}

const (
	OpcodeAdd         = 1
	OpcodeMultiply    = 2
	OpcodeTerminate   = 99
	InstructionLength = 4
)

func ComputeIntcode(input string) string {
	ints := []int{}
	nums := strings.Split(input, ",")
	for _, s := range nums {
		n, err := strconv.Atoi(s)
		try(err)
		ints = append(ints, n)
	}
	out := computeIntcode(ints, 0)
	strs := []string{}
	for _, n := range out {
		s := strconv.Itoa(n)
		strs = append(strs, s)
	}
	return strings.Join(strs, ",")
}

func computeIntcode(input []int, pos int) []int {
	opcode := input[pos]
	if opcode == OpcodeTerminate {
		return input
	}

	out := make([]int, len(input))
	copy(out, input)

	operand1Reg := out[pos+1]
	operand2Reg := out[pos+2]
	operand1 := out[operand1Reg]
	operand2 := out[operand2Reg]
	outReg := out[pos+3]

	if opcode == OpcodeAdd {
		out[outReg] = operand1 + operand2
	} else if opcode == OpcodeMultiply {
		out[outReg] = operand1 * operand2
	} else {
		panic(fmt.Sprintf("Unexpected opcode: %d from position: %d", opcode, pos))
	}

	return computeIntcode(out, pos+InstructionLength)
}

func Day2Part1(input string) string {
	return ComputeIntcode(input)
}
