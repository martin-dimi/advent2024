package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputFileArg := os.Args[1]

	// Read input file into string
	b, _ := os.ReadFile(inputFileArg)

	part1(b)
	part2(b)
}

func part1(input []byte) {
	r := `mul\(\d{1,6},\d{1,6}\)`

	// Find all matches
	result := 0
	matches := regexp.MustCompile(r).FindAll(input, -1)
	for _, match := range matches {
		mul := string(match)
		mul = mul[4 : len(mul)-1]
		values := strings.Split(mul, ",")

		a, _ := strconv.Atoi(values[0])
		b, _ := strconv.Atoi(values[1])

		result += a * b
	}

	fmt.Printf("Part 1: Result %v\n", result)
}

func part2(input []byte) {
	r := `mul\(\d{1,6},\d{1,6}\)|don't\(\)|do\(\)`

	// Find all matches
	result := 0
	enable := true
	matches := regexp.MustCompile(r).FindAll(input, -1)
	for _, match := range matches {
		val := string(match)

		if val == "do()" {
			enable = true
			continue
		}
		if val == "don't()" {
			enable = false
			continue
		}

		if enable == false {
			continue
		}

		val = val[4 : len(val)-1]
		values := strings.Split(val, ",")

		a, _ := strconv.Atoi(values[0])
		b, _ := strconv.Atoi(values[1])

		result += a * b
	}

	fmt.Printf("Part 1: Result %v\n", result)
}
