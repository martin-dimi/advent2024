package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputFileArg := os.Args[1]

	file, err := os.Open(inputFileArg)
	if err != nil {
		panic(err)
	}

	left := make([]int, 0)
	right := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		words := strings.Fields(line)
		l, _ := strconv.Atoi(words[0])
		r, _ := strconv.Atoi(words[1])

		left = append(left, l)
		right = append(right, r)
	}

	part1(left, right)
	part2(left, right)
}

func part1(leftList, rightList []int) {
	// sort
	sort.Ints(leftList)
	sort.Ints(rightList)

	distance := 0
	for i := 0; i < len(leftList); i++ {
		distance += abs(rightList[i] - leftList[i])
	}

	fmt.Printf("Part 1: Distance = %d\n", distance)
}

func part2(leftList, rightList []int) {
	rightMap := make(map[int]int)
	for _, r := range rightList {
		rightMap[r] = rightMap[r] + 1
	}

	similarityScore := 0
	for _, l := range leftList {
		similarityScore += l * rightMap[l]
	}

	fmt.Printf("Part 2: Similarity = %d\n", similarityScore)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
