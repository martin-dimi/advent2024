package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFileArg := os.Args[1]

	part1(inputFileArg)
	// part2(inputFileArg)
}

func part1(inputFileArg string) {
	// Read input file into string
	b, _ := os.ReadFile(inputFileArg)
	input := string(b)
	rows := strings.Split(input, "\n")
	rows = rows[:len(rows)-1]
	cols := len(rows[0])

	// Wrap the input in . border
	rows = append([]string{strings.Repeat(".", cols)}, rows...)
	rows = append(rows, strings.Repeat(".", cols))
	for i, row := range rows {
		rows[i] = "." + row + "."
	}

	text := strings.Join(rows, "")
	occurrences := 0
	for i := 0; i < len(text); i++ {
		if text[i] == 'X' {
			occurrences += checkWord(text, i, len(rows[0]))
		}
	}

	fmt.Printf("Part 1: Occurrences %v\n", occurrences)
}

// The word can be horizontal, vertical, diagonal and even in reverse.
func checkWord(text string, startIndex int, totalCols int) int {
	xDirs := []int{-1, 0, 1}
	yDirs := []int{-1, 0, 1}

	matches := 0
	for _, xDir := range xDirs {
		for _, yDir := range yDirs {
			if xDir == 0 && yDir == 0 {
				continue
			}

			found := true
			match := "XMAS"
			for i := 0; i < 4; i++ {
				xOffset := i * xDir
				yOffset := i * totalCols * yDir
				idx := startIndex + xOffset + yOffset

				if idx < 0 || idx >= len(text) {
					found = false
					break
				}

				if text[idx] != match[i] {
					found = false
					break
				}
			}

			if found {
				matches++
			}
		}
	}

	return matches
}

func part2(inputFileArg string) {
	// Read input file into string
	b, _ := os.ReadFile(inputFileArg)
	input := string(b)
	rows := strings.Split(input, "\n")
	cols := len(rows[0])
	text := strings.Join(rows, "")

	occurrences := 0
	for i := 0; i < len(text); i++ {
		if text[i] == 'A' && checkXWord(text, i, cols, len(rows)-1) {
			occurrences++
		}
	}

	fmt.Printf("Part 2: Occurrences %v\n", occurrences)
}
func checkXWord(text string, startIndex int, totalCols int, totalRows int) bool {
	xCord := startIndex % totalCols
	yCord := startIndex / totalCols

	// Nothing on the edge.
	if xCord == 0 || xCord == totalCols-1 ||
		yCord == 0 || yCord == totalRows-1 {
		return false
	}

	topLeft := startIndex - totalCols - 1
	bottomRight := startIndex + totalCols + 1

	topRight := startIndex - totalCols + 1
	bottomLeft := startIndex + totalCols - 1

	return areMandS(text[topLeft], text[bottomRight]) &&
		areMandS(text[topRight], text[bottomLeft])
}

func areMandS(l1, l2 byte) bool {
	return (l1 == 'M' && l2 == 'S') || (l1 == 'S' && l2 == 'M')
}
