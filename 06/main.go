package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3
)

func main() {
	intputFileArg := os.Args[1]

	part1(intputFileArg)
	part2(intputFileArg)
}

func part2(intputFileArg string) {
	f, _ := os.ReadFile(intputFileArg)
	field := string(f)
	r := strings.Split(field, "\n")
	r = r[:len(r)-1]
	rows := len(r)
	cols := len(r[0])

	field = strings.Join(r, "")
	startLocation := strings.Index(field, "^")
	startPosition := UP
	guardStartLocation = startLocation
	// field = strings.Replace(field, "^", "|", 1)

	lookLoops(field, rows, cols, startLocation, startPosition, make(map[int][]int), 1)

	fmt.Printf("Part 2: Crosses %v\n", len(crosses))
}

var (
	guardStartLocation = 0
	crosses            = map[int]struct{}{}
	XVector            = map[int]int{
		RIGHT: 1,
		LEFT:  -1,
	}
	YVector = map[int]int{
		UP:   -1,
		DOWN: 1,
	}
	PVector = map[int]string{
		LEFT:  "-",
		RIGHT: "-",
		UP:    "|",
		DOWN:  "|",
	}
)

func lookLoops(field string, rows, cols int, location int, direction int, locToDirs map[int][]int, extra int) (string, bool) {
	currentX := location % cols
	currentY := location / cols

	for true {
		nextX := currentX + XVector[direction]
		nextY := currentY + YVector[direction]
		currentLoc := currentY*cols + currentX
		nextLoc := nextY*cols + nextX
		nextDirection := (direction + 1) % 4

		if nextY < 0 || nextY >= rows || nextX < 0 || nextX >= cols {
			return field, false
		}

		if slices.Contains(locToDirs[currentLoc], direction) {
			return field, true
		}

		locToDirs[currentLoc] = append(locToDirs[currentLoc], direction)
		currentVal := string(field[currentLoc])
		newVal := PVector[direction]
		if newVal+currentVal == "-|" || newVal+currentVal == "|-" || currentVal == "+" {
			newVal = "+"
		}
		field = field[:currentLoc] + newVal + field[currentLoc+1:]

		if _, ok := crosses[nextLoc]; !ok && extra > 0 && nextLoc != guardStartLocation && field[nextLoc] == '.' {
			temp := field[:nextLoc] + "0" + field[nextLoc+1:]
			ops, ok := lookLoops(temp, rows, cols, currentLoc, nextDirection, deepCloneMap(locToDirs), extra-1)
			if ok && len(ops) > 0 {
				crosses[nextLoc] = struct{}{}
				// fmt.Printf(strings.Join(Chunks(ops, cols), "\n") + "\n\n")
			}
		}

		if field[nextLoc] == '#' || field[nextLoc] == '0' {
			field = field[:currentLoc] + "+" + field[currentLoc+1:]
			return lookLoops(field, rows, cols, currentLoc, nextDirection, locToDirs, extra)
		}

		currentX = nextX
		currentY = nextY
	}

	return field, false
}

func part1(intputFileArg string) {
	f, _ := os.ReadFile(intputFileArg)
	field := string(f)
	r := strings.Split(field, "\n")
	r = r[:len(r)-1]
	rows := len(r)
	cols := len(r[0])
	// fmt.Println(field)

	field = strings.Join(r, "")
	startLocation := strings.Index(field, "^")
	startPosition := UP
	field = strings.Replace(field, "^", "X", 1)

	field = solveRoute(field, rows, cols, startLocation, startPosition)

	answer := strings.Count(field, "X")
	fmt.Printf("Part 1: Answer %v\n", answer)
	// fmt.Println(strings.Join(Chunks(field, cols), "\n"))
}

func solveRoute(field string, rows, cols int, location int, direction int) string {
	XVector := map[int]int{
		RIGHT: 1,
		LEFT:  -1,
	}
	YVector := map[int]int{
		UP:   -1,
		DOWN: 1,
	}

	currentX := location % cols
	currentY := location / cols

	for true {
		nextX := currentX + XVector[direction]
		nextY := currentY + YVector[direction]
		if nextY < 0 || nextY >= rows || nextX < 0 || nextX >= cols {
			return field
		}

		currentLoc := currentY*cols + currentX
		nextLoc := nextY*cols + nextX
		if field[nextLoc] == '#' {
			nextDirection := (direction + 1) % 4
			return solveRoute(field, rows, cols, currentLoc, nextDirection)
		}

		currentX = nextX
		currentY = nextY
		field = field[:nextLoc] + "X" + field[nextLoc+1:]
	}

	return field
}

func Chunks(s string, chunkSize int) []string {
	if len(s) == 0 {
		return nil
	}
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == chunkSize {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}

func deepCloneMap(m map[int][]int) map[int][]int {
	newMap := make(map[int][]int)
	for k, v := range m {
		newMap[k] = append([]int{}, v...)
	}
	return newMap
}
