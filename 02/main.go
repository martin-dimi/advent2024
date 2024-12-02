package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFileArg := os.Args[1]

	part1(inputFileArg)
	part2(inputFileArg)
}

func part1(inputFileArg string) {
	file, err := os.Open(inputFileArg)
	if err != nil {
		panic(err)
	}

	saveReports := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		inc := 1
		prev := 0
		valid := true
		for i, v := range strings.Fields(line) {
			number, _ := strconv.Atoi(v)

			if i == 0 {
				prev = number
				continue
			}

			if i == 1 {
				if number < prev {
					inc = -1
				} else {
					inc = 1
				}
			}

			distance := abs(number - prev)
			if distance < 1 || distance > 3 {
				valid = false
				break
			}

			if i > 1 && (number-prev)*inc < 0 {
				valid = false
				break
			}

			prev = number
		}

		if valid {
			saveReports++
		}
	}

	fmt.Printf("Part 1: Number of valid reports = %d\n", saveReports)
}

func part2(inputFileArg string) {
	file, err := os.Open(inputFileArg)
	if err != nil {
		panic(err)
	}

	saveReports := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if isReportValid(strings.Fields(line), nil, 2) {
			saveReports++
		}
	}

	fmt.Printf("Part 2: Number of valid reports = %d\n", saveReports)
}

func isReportValid(report []string, shouldIncrement *bool, lives int) bool {
	if lives == 0 {
		return false
	}
	if len(report) < 2 {
		return true
	}

	l, _ := strconv.Atoi(report[0])
	r, _ := strconv.Atoi(report[1])

	inc := l < r
	distance := abs(l - r)

	// consider: x, y, z

	current := distance >= 1 && distance <= 3 && (shouldIncrement == nil || inc == *shouldIncrement)
	if !current {
		// if current is invalid, there's 2 cases dropping x (can change direction) or y
		if shouldIncrement == nil {
			return isReportValid(splice(report, 0), nil, lives-1) || isReportValid(splice(report, 1), nil, lives-1)
		}

		//
		return isReportValid(splice(report, 1), shouldIncrement, lives-1)
	}

	if !isReportValid(report[1:], &inc, lives) {
		// If the future fails, try dropping either the left or right entries.
		return isReportValid(splice(report, 0), shouldIncrement, lives-1) || isReportValid(splice(report, 1), shouldIncrement, lives-1)
	}

	return true
}

func splice(s []string, i int) []string {
	new := make([]string, 0, len(s)-1)
	new = append(new, s[:i]...)
	new = append(new, s[i+1:]...)
	return new
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
