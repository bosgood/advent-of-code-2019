package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func try(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("inputs/day14_part1.txt")
	try(err)
	b, err := ioutil.ReadAll(f)
	try(err)
	fmt.Println(Day14Part1(string(b)))
}

const (
	PositionX    = 0
	PositionY    = 1
	PositionTile = 2
)

func Day14Part1(input string) string {
	tokens := strings.Split(input, ",")
	// tiles := map[string]bool{}
	tuple := ""
	lines := []string{}
	for i, c := range tokens {
		pos := i % 3
		// if pos == PositionTile {
		// 	tiles[c] = true
		// }
		switch pos {
		case PositionX:
			tuple = c
			break
		case PositionY:
			tuple += "," + c
			break
		case PositionTile:
			tuple += "," + c
			lines = append(lines, fmt.Sprintf("(%s)", tuple))
			break
		}
	}
	return strings.Join(lines, "\n")
}
