package main

import (
	"fmt"
	"strconv"
	"strings"
)

func try(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	input := "1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,5,19,23,2,10,23,27,1,27,5,31,2,9,31,35,1,35,5,39,2,6,39,43,1,43,5,47,2,47,10,51,2,51,6,55,1,5,55,59,2,10,59,63,1,63,6,67,2,67,6,71,1,71,5,75,1,13,75,79,1,6,79,83,2,83,13,87,1,87,6,91,1,10,91,95,1,95,9,99,2,99,13,103,1,103,6,107,2,107,6,111,1,111,2,115,1,115,13,0,99,2,0,14,0"
	output := 19690720

	fmt.Println(Part1(input) + "\n")
	fmt.Println(Part2(input, output) + "\n")
}

const (
	OpcodeAdd         = 1
	OpcodeMultiply    = 2
	OpcodeTerminate   = 99
	InstructionLength = 4
)

func ComputeIntcode(input string) string {
	ints := parse(input)
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

func Part1(input string) string {
	return ComputeIntcode(input)
}

func Part2(input string, output int) string {
	test := parse(input)
	limit := len(test) - 1
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			ints := parse(input)
			ints[1] = i
			ints[2] = j
			out := computeIntcode(ints, 0)
			if out[0] == output {
				return fmt.Sprintf("noun=%d, verb=%d, answer=%d", i, j, 100*i+j)
			}
		}
	}
	return "not found"
}

func parse(input string) []int {
	ints := []int{}
	nums := strings.Split(input, ",")
	for _, s := range nums {
		n, err := strconv.Atoi(s)
		try(err)
		ints = append(ints, n)
	}
	return ints
}
