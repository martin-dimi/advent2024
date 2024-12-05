package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fileName := os.Args[1]
	inputBytes, _ := os.ReadFile(fileName)

	lines := strings.Split(string(inputBytes), "\n")

	readingRules := true
	rules := make(map[int][]int)
	updates := [][]int{}
	for _, line := range lines {
		if len(line) == 0 {
			readingRules = false
			continue
		}

		if readingRules {
			nums := strings.Split(line, "|")
			l, _ := strconv.Atoi(nums[0])
			r, _ := strconv.Atoi(nums[1])

			rules[l] = append(rules[l], r)
		} else {
			nums := strings.Split(line, ",")
			update := []int{}
			for _, num := range nums {
				n, _ := strconv.Atoi(num)
				update = append(update, n)
			}

			updates = append(updates, update)
		}
	}

	part1(rules, updates)
	part2(rules, updates)
}

func part2(rules map[int][]int, updates [][]int) {
	answer := 0
	for _, update := range updates {
		valid := true

	out:
		for i := 0; i < len(update)-1; i++ {
			l := update[i]

			for j := i + 1; j < len(update); j++ {
				r := update[j]
				validLefts := rules[r]

				if slices.Contains(validLefts, l) {
					valid = false
					break out
				}
			}
		}

		if valid {
			continue
		}

		for i := 0; i < len(update)-1; i++ {
			l := update[i]

			swapped := false
			for j := i + 1; j < len(update); j++ {
				r := update[j]
				validLefts := rules[r]

				if slices.Contains(validLefts, l) {
					// fmt.Println("Swapping", l, r)
					// swap the two elements and restart the for loops
					update[i] = r
					update[j] = l
					i = -1
					j = -1
					swapped = true
					break
				}
			}

			if swapped {
				continue
			}

		}

		// fmt.Println(update)

		middle := getMiddleElementOfSlice(update)
		answer += middle

	}

	fmt.Printf("Part 2 Answer: %v\n", answer)
}

func part1(rules map[int][]int, updates [][]int) {
	answer := 0
	for _, update := range updates {
		// 75,47,61,53,29

		valid := true

	out:
		for i := 0; i < len(update)-1; i++ {
			l := update[i]

			for j := i + 1; j < len(update); j++ {
				r := update[j]
				validLefts := rules[r]

				if slices.Contains(validLefts, l) {
					valid = false
					break out
				}
			}
		}

		if valid {
			middle := getMiddleElementOfSlice(update)
			answer += middle
		}
	}

	fmt.Printf("Part 1 Answer: %v\n", answer)
}

func getMiddleElementOfSlice(s []int) int {
	if len(s) == 0 {
		return -1
	}
	if len(s) == 1 {
		return s[0]
	}
	return s[len(s)/2]
}
