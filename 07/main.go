package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Eq struct {
	total int
	nums  []int
}

func (e Eq) Valid() bool {
	return e.total == 0 && len(e.nums) == 0
}

func (e Eq) String() string {
	return fmt.Sprintf("%v: %v", e.total, e.nums)
}

func main() {
	fileName := os.Args[1]
	b, _ := os.ReadFile(fileName)
	lines := strings.Split(string(b), "\n")
	lines = lines[:len(lines)-1]

	equations := []Eq{}
	for _, l := range lines {
		parts := strings.Split(l, ":")

		total, _ := strconv.Atoi(parts[0])
		nums := []int{}

		numS := strings.Fields(parts[1])
		for _, n := range numS {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}

		equations = append(equations, Eq{total, nums})
	}

	part1(equations)
	part2(equations)
}

func part2(equations []Eq) {
	answer := 0
	for _, eq := range equations {
		if validEq2(eq) {
			fmt.Println(eq)
			answer += eq.total
		}
	}

	fmt.Printf("\nPart 2 Answer = %v\n", answer)
}

func validEq2(eq Eq) bool {
	if len(eq.nums) == 1 {
		return eq.nums[0] == eq.total
	}

	l := eq.nums[0]
	r := eq.nums[1]

	// pluses
	if validEq2(Eq{total: eq.total, nums: append([]int{l + r}, eq.nums[2:]...)}) {
		return true
	}
	if validEq2(Eq{total: eq.total, nums: append([]int{l * r}, eq.nums[2:]...)}) {
		return true
	}
	if validEq2(Eq{total: eq.total, nums: append([]int{concat(l, r)}, eq.nums[2:]...)}) {
		return true
	}

	return false
}

func part1(equations []Eq) {
	answer := 0
	for _, eq := range equations {
		if validEq1(eq) {
			// fmt.Println(eq)
			answer += eq.total
		}
	}

	fmt.Printf("\nPart 1 Answer = %v\n", answer)
}

func validEq1(eq Eq) bool {
	if len(eq.nums) == 1 {
		return eq.nums[0] == eq.total
	}

	l := eq.nums[0]
	r := eq.nums[1]

	// pluses
	if validEq1(Eq{total: eq.total, nums: append([]int{l + r}, eq.nums[2:]...)}) {
		return true
	}
	if validEq1(Eq{total: eq.total, nums: append([]int{l * r}, eq.nums[2:]...)}) {
		return true
	}

	return false
}

func sum(nums []int) int {
	v := 0
	for _, n := range nums {
		v += n
	}
	return v
}

func times(nums []int) int {
	v := 1
	for _, n := range nums {
		v *= n
	}
	return v
}

func concat(n1, n2 int) int {
	v, _ := strconv.Atoi(fmt.Sprintf("%v%v", n1, n2))
	return v
}
