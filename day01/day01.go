package day01

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/szuduko/aoc-go/util"
)

func Part1() {
	fmt.Println("Part 1")
	lines := util.ReadInput("day01/day01-input.txt")
	
	sum := 0
	re := regexp.MustCompile("[0-9]")

	for _, line := range lines {
		digits := re.FindAllString(line, -1)
		first := string(digits[0])
		last := string(digits[len(digits) - 1])

		number, err := strconv.Atoi(string(first) + string(last))
		
		util.Check(err)

		sum += number
	}

	fmt.Println(sum)
}

func Part2() {
	fmt.Println("Part 2")
}